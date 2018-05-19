#!/bin/bash

# error exit codes
readonly ERROR_PARSING_ACTION=192

readonly SERVER_COMMAND='SimpleHTTPServer 9090'

# store the global exit status
EXIT_STATUS=0

# Defines the main script flow
function main()
{
	# Print usage info if no arguments were provided
	if [[ "$#" -eq 0 ]]; then
		print_usage "$@"
	fi
	local ACTION=$1

	# Parse and check the given action
	local PARSED_ACTION
	PARSED_ACTION=$(parse_action "${ACTION}")
	validate_no_errors "${PARSED_ACTION}" "${ERROR_PARSING_ACTION}"

	case ${PARSED_ACTION} in
	status)
		go_status
		;;
	start)
		go_start
		;;
	stop)
		go_stop
		;;
	*)
		print_usage "$@"
		;;
	esac

	return "${EXIT_STATUS}"
}


# Prints the script's usage info
function print_usage()
{
	local COMMAND_PARAM=$1

	echo "This script provides a CLI for the project, allowing it to be run via simple English statements, and pre-validating its arguments."
	echo ""
	echo "Usage:"
	echo "  ./go <action>"
	echo ""
	get_supported_actions
	echo "Examples:"
	echo "  ./go start -> Start the webserver for the slides"

	if [ "${COMMAND_PARAM}" = "" ]; then
		exit ${EXIT_STATUS}
	else
		exit ${ERROR_PARSING_ACTION}
	fi
}


function print_divider()
{
	# Standard 80-character divider
	echo "--------------------------------------------------------------------------------"
}


# Returns a formatted list of actions supported by the ./go script
function get_supported_actions()
{
	local SUPPORTED_ACTIONS="  Supported actions:\\n"

	# Add each action, one line per action
	SUPPORTED_ACTIONS="${SUPPORTED_ACTIONS}    - status\\n"
	SUPPORTED_ACTIONS="${SUPPORTED_ACTIONS}    - start\\n"
	SUPPORTED_ACTIONS="${SUPPORTED_ACTIONS}    - stop\\n"

	echo -e "${SUPPORTED_ACTIONS}"
}


function get_server_pid()
{
	local SERVER_PID
	SERVER_PID=$(ps -ef | grep "${SERVER_COMMAND}" | grep -v grep | tr -s " " | cut -d" " -f3)
	echo "${SERVER_PID}"
}

function go_status()
{
	local SERVER_PID
	SERVER_PID=$(get_server_pid)

	if [[ "${SERVER_PID}" == "" ]]; then
		echo "Slide server is stopped"
	else
		echo "Slide server is running"
	fi
}


function go_start()
{
	python -m ${SERVER_COMMAND} &
}


function go_stop()
{
	local SERVER_PID
	SERVER_PID=$(get_server_pid)

	if [[ "${SERVER_PID}" != "" ]]; then
		kill "${SERVER_PID}"
	fi
}


# Simple parsing that the given action is supported
# Returns either the validated action or an error string
function parse_action()
{
	local GIVEN_ACTION=$1
	local ERROR="ERROR: Unsupported action: ${GIVEN_ACTION}"

	# Preset the result to a default of the error - only success will overwrite this
	local RESULT=${ERROR}

	# Iterating below will populate the list of supported actions that can be chosen
	local SUPPORTED_ACTIONS
	SUPPORTED_ACTIONS="$(get_supported_actions)"

	if [[ "${GIVEN_ACTION}" != "" ]]; then
		# Go through each of the lines in SUPPORTED_ACTIONS
		while read -r CURRENT_ACTION; do
			# For any that are formatted as an action, strip them down to the bare action name
			# E.g. '    - check' becomes 'check'
			local ACTION_SHORTNAME
			ACTION_SHORTNAME=$(echo "${CURRENT_ACTION}" | sed -n 's/.*- \(.*\).*/\1/p')

			# If we find the droid we're looking for ;-)
			if [[ "${ACTION_SHORTNAME}" = "${GIVEN_ACTION}" ]]; then
				RESULT="${ACTION_SHORTNAME}"
			fi
		done <<< "${SUPPORTED_ACTIONS}"
	fi

	# If we haven't overwritten the result, it will still be preset to the error
	if [[ "${RESULT}" = "${ERROR}" ]]; then
		echo -e "${ERROR}\\n${SUPPORTED_ACTIONS}\\n"
		exit ${ERROR_PARSING_ACTION}
	else
		echo "${RESULT}"
		exit ${EXIT_STATUS}
	fi
}


# Checks if the given parsed input failed parsing and returned an error
# If an error was found, the shell execution terminates with the given exit code
function validate_no_errors()
{
	local PARSED_INPUT=$1
	local EXIT_CODE=$2

	if [[ $PARSED_INPUT =~ ^ERROR.* ]]; then
		echo "${PARSED_INPUT}"
		exit "${EXIT_CODE}"
	fi
}


# Execute the main program flow with all arguments
main "$@"
