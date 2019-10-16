---
layout: post
title:  "Infrastructure as Code"
date:   2019-10-16 23:32:02 +0100
categories: iac infrastructure code
---
Infrastructure as Code (IaC) is a term used to describe a collection of practices for managing IT resources via programmable interfaces, allowing infrastructure to be defined in code, rather than created through a series of human activity. Typically, infrastructure resources might include various types of network, storage, and compute components, alongside resources relating to identity and permissions. The need for IaC arose out of the growing number of problems encountered trying to manage increasingly complex IT with traditional techniques. The demands of modern business typically bring the need to scale - both in volume and frequency, whilst maintaining quality and consistency. This typically results in compromising stability to meet a relentless need for business change.

Defining infrastructure as code goes way beyond scripting or config management - it enables engineers to leverage many of the well-established good practices that software engineers have been using for decades. Not just fundamental programming constructs like variables, conditionals and loops, but higher level abstractions such as design patterns, and highly effective practices such as version control, dependency management, automated testing, continuous integration, and continuous delivery.

Centiq leverage IaC to removes the complexity of SAP landscape management - simplifying operational activities into automated workflows such as _build system_, _stop system_, _start system_, or _failover system_; and keeping this consistent across all system types whether they are clustered, have DR capabilities, or even across different SAP products such as S/4HANA or BW/4HANA. Even OS-specific differences are encapsulated within the workflow, so the user experience is identical nomatter what OS is your preference.

In future posts, I'll explore some of the ways we incorporate IaC practices into our projects and services at [Centiq][centiq].

[centiq]: https://centiq.co.uk
