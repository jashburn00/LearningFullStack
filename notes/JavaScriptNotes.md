# JavaScript Notes
## Contents
<!-- no toc -->
- [General Use Cases](#general-use-cases)

## General Use Cases
- anything that isn't static information
- JS responds to browser events (button click, key press, etc.)
- can be embedded in html file using the "< script>" tag

## Things I Don't Already Know
- __default parameter value__ can be specified using the syntax `function hello(name = 'Scudd'){`
- __template strings__ (aka template literals) use backticks instead of single/dbl quotes
    - ex: ` `hello ${name}! This is a string!`  `
- __let__ vs __var__: 
    - let is only block scoped, while var is function scoped
- __function expressions__ are functions declared as variables (e.g. const doMath = function() )
    - these use anonymous functions and so they are not hoisted (regular JS functions are hoisted)
    - can use the "arrow function syntax" e.g. `const getArea = (l, w) => {`
        - people love this shit apparently because there are severals ways to declare:
            - const getHelp = () => {
            - const getHelp = name => {
            - const getHelp = (name, age) => {
            - const getHelp = age => age+2; (single line only)
- identically named variables will be used in closest scope order (nearest declaration used)
- there are more bulitin functions than I know/use (array.splice is cool)
- javascript is __pass by value__ for primitive values (bool, int, etc.) but is __pass by reference__ for objects

