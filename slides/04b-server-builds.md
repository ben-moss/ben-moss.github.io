### Server Builds

<!-- .slide: data-background="../assets/examples/build/snowflakes.png" -->

---

<!-- .slide: data-background="../assets/examples/build/snowflakes.png" -->

### Problems: Snowflake Servers

Manually-built servers are fragile snowflakes

- Changes applied from variety of sources
  - varying visibility
  - difficult to sustain control 
- No capability to safely test changes <!-- .element: class="fragment" -->
  - changes break live service
- Cannot reliably provision environments <!-- .element: class="fragment" -->
  - new ones or replacements
- Cannot reliably test on environments <!-- .element: class="fragment" -->
  - environment is not representative


<!-- <object type="image/svg+xml" data="assets/examples/build/snowflakes-server-build1.svg#grey">
	<param id="purple" class="fragment" data-fragment-index="1">
	<param id="green" class="fragment" data-fragment-index="2">
	<param id="red" class="fragment" data-fragment-index="3">
	<param id="blue" class="fragment" data-fragment-index="4">
</object>
 -->

---

<!-- .slide: data-background="../assets/examples/build/snowflakes.png" -->

### Solution: Configuration Management

<img data-src="assets/examples/build/config-management.png">

---

<!-- .slide: data-background="../assets/examples/build/snowflakes.png" -->

### Solution: Configuration Management Examples

<img data-src="assets/examples/build/config-management-examples.png">

---

<!-- .slide: data-background="../assets/examples/build/snowflakes.png" -->

### Business Outcomes Delivered

- Reduced risk by increasing consistency through the lifecycle <!-- .element: class="fragment" -->
- Reduce cycle time to build/rebuild environments <!-- .element: class="fragment" -->
- Inspect the state without having access to the environments themselves <!-- .element: class="fragment" -->
