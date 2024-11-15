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
- __loop syntaxes__:
    - javascript simplified for loop syntax: `for(const b of balls)`
    - other syntax for iterating all properties of objects: `for (let crewMate in spaceship.crew)`
        - this does not iterate objects themselves, so `crewMate.name` does not work 
        - you would have to do `ship.crew[crewMate].name`
- methods are functions which are stored in objects as data members, not the other way around
    - the syntax for this in JavaScript is: `doSomething: function (){}` or `doSomething() {//code}

## Objects: JS Edition
- JS Object syntax and idiosyncracies
    - __object literals__ can be created with `let balls = {}`
    - objects use key value pairs but can omit quotations for the keys
    - accessing object data members can also be done with array syntax e.g. `balls[texture]`
    - assigning a value to an object property that doesn't exist will CREATE A NEW PROPERTY (wtf)
    - object properties can be deleted like so: `delete myObj.name`
    - objects declared as const are still mutable, just not reassignable
- this
    - required for referencing data in non-static objects (unlike Java, my tru love)
    - declaring methods with arrow syntax keeps them anonymous, and they will always have `this` refer to the JS global object
- privacy
    - does not exist in JS, so devs prefix members with _ to indicate private
        - "I been around the world, I'm not impressed."
- getters & setters
    - syntax is `get fullName() {` and `set fullName(name){}`
    - this allows the getter/setter functions to behave as normal property assignments
        - e.g. `myClass.name = 10;`
    - cannot share name with any data member (except for the corresponding getter/setter)
- __destructuring__:
    - object creation: `const factory = (name, age) => { return{name, age}}`
        - basically you don't have to write the identifier twice. cool.
    - object destructuring assignment: `const { hoof, joints } = deer` (where deer is an object with properties hoof and joints)
        - *hot tip* this can also be done with arrays: `const [first, second] = myArray`
            - `const [first, second, , fourth] = myArray`
- generic Object methods
    - Object.keys(myObj) returns the names of the keys in an object
    - Object.entries(myObj) returns the key value pairs of an object as a 2D array
    - Object.assign(targetObj, basisObj, additionalPairs) 
        - this one assigns to the target (which can be set to {} for a new object) a new object using a copy of the basis object parameter and can be given additional key value pairs optionally
### functions
- functions are objects in JavaScript, and have properties such as name 
- __high order functions__ are functions that use functions as parameter and/or return type
- syntax for passing in functions as arguments: pass in only the identifier, with no parentheses
- e.g. `cleanData(useData, data);`
- examples of high order functions:
    - array builtin methods like `filter`:
        -`const newList = words.filter(word => {return word.length == 3;});`











make it shake make it clap -kodak black
