# JavaScript Notes
## Contents
<!-- no toc -->
- [General Use Cases](#general-use-cases)
- [Things I Don't Already Know](#things-i-dont-already-know)
- [Objects: JS Edition](#objects-js-edition)
- [Website Interactivity](#website-interactivity)

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

## Website Interactivity 
### Scripts
- included using the `<scrit >` tag (must be closed)
  - *src* property allows url to .js file
- scripts are loaded and executed when reached by the HTML parser, potentially slowing load time
- scripts are read in order of appearance, so order matters for dependencies
- `defer` and `async` are used to combat this
  - scripts with the defer attribute (no value) will load but not execute until the rest of the HTML is parsed
    - this is useful when the script needs the DOM to finish parsing first
  - the async attribute (no value) will cause the script to be loaded in the background as parsing continues
    - useful for scripts that do not depend on other scripts, as it optimizes load time
### The DOM
- tree-like model that allows JS to access HTML elements
- the `document` node is the root node of any DOM
- elements and their properties can both be accessed through the DOM
- you can append a child element in the DOM by setting the .innerHTML of an element to a valid HTML element
 
### Events
- events fire on objects in the DOM
- DOM elements can be given event listening functionality using `elt.addEventListener('event', function)`
    - this way of adding event listeners allows for more than one to be added
    - note the syntax of no parentheses in function name when adding event listeners 
        - the browser automatically passes the Event object to the event handler function
- event listeners can also be done using `elt.onevent = myfunc`
- removed with `elt.removeEventListener('event', func)`
- event properties:
    - .target - the element which the event was registered to
    - .type - the type of event
    - .timeStamp - time in milliseconds since the document loaded and the event was triggered

### Form Validation
- term for ensuring bad data is not committed to a server or database
- e.g. matching regex or testing values upon user submission
#### Client-side validation
- Regex
    - use square brackets to include a list of characters, or ^ to denote any
        - followed with asterisk to indicate any amount of this pattern
    - use \d{X} to denote X numeric digits
    - use literal characters as well 
- can also use libraries
#### server side/back-end validation
- all data should initially be treated as invalid

### Forms
- must be created with the location where data will be sent and the type of request 
    - ex: `<f orm action="/example.html" method="post"`
- use < input> values
    - ex: < input type="text" name="username" value="default text">
        - type text can use the 'minlength' and 'maxlength' attributes
    - "password" type value to obscure text automatically 
    - "number" type allows limited symbols (-, +, .) 
        - add the `step` property to include arrow buttons, and the increment amount as its value
        - can use the 'min' and 'max' attributes
    - "range" type accompanied by a `min` and `max` value can also use step
        - this creates a slider
    - "checkbox" 
        - use the name and value attributes like other elements (but be wary that they do not render)
            - the value property is what gets submitted in data; if unchecked, nothing is added
        - group them using the name attribute (same name)
    - "radio"
        - you know how it go
    - dropdown lists
        - created with a < select> element with child < option> elements
        - only requires the "value" attribute
    - datalist
        - used in place of < select> (instead use < datalist>) 
        - uses child < options> but they are now searchable
        - used as the backend of a "text" type input when the input uses the attribute list="mylist"
            - where mylist is the id of your datalist
    - textarea (not a type of input object)
        - has attributes rows="2" and cols="30" for sizing 
        - set default text between the closing and opening tags 
- use the 'required' attribute to make fields mandatory
- many types of input can use the 'pattern' attribute, which allows a regex
- when form is submitted, input key value pairs are sent in the http request
    - ex: "username=jeff100"
- use labels for input fields
    - ex: `< label for="inputfieldID">text</ label>`
- submitting forms
    - simply use input of type "submit", and use the value attribute to set button text

#### Using the form in JavaScript
- assuming the form has been sent to 'page.html':
1. page.html must use a JS source file script
2. in the linked JS file, collect the values like so:
    > const words = new URLSearchParams(window.location.search);
3. now words contains the key value pairs from the request, and can be accessed like so:
    > words.get('animal-1')
4. now you are free to use them (populate a page, etc.)

## Classes 
- example class: `
    class Dog {
        constructor(name) {
            this._name = name;
            this._age = 0;
        }

        get age(){
            return this._age;
        }

        birthday(){
            this._age++;
            return this._age;
        }
    }

    let sparky = new Dog('sparky');
`
- JS classes always use `this` to refer to an instance (such as itself)
- uses `new` keyword to create instances
- note the use of the get and set keywords
    - use sparky.age as if it was a public data member in Java
- note the use of the _ prefix to indicate private data members

### Inheritance
- uses the `extends` keyword 
- unlike Java, the super call is explict and must be done before any reference to the instance
    `
    class WeenerDog extends Dog{
        constructor(name, length){
            super(name);
            this._length = length;
    }
    `
- static methods are delcared by adding the keyword `static` before the function name


## Modules
- often used interchangeably with "files"
- used to create modular systems and separate concerns
### The Basics
- importing/exporting code between modules uses different syntax for the two main runtimes
    - Node.js runtime 
        - uses the `module.exports` and `require()` syntax
            - to export things, create a property of the object `module.exports`
                - e.g. `module.exports.sayHello = function sayHello(){ ...`
                - e.g. `module.export.sayHello = sayHello;
            - to import it elsewhere, use `require()` with the file path of the desired module
                - e.g. `const myModule = require('./modules/myModule');`
                - you can specify specific parts of an imported module like so:
                    - `const {moduleFunction} = require('./myModule');`
            - imported functions do not need to be prepended with the module name
        - uses `process` instead of browser environments like window and document
            - can use `process.argv[2]` to access parameters from command line runs
        - has filesystem library `fs`
            - create a new file: `fs.appendFileSync(path, str)`
                - if a file at path exists, it will append str to it; otherwise it makes a new file with str as content
            - read a file: `fs.readFileSync(path)` returns contents of file found
    - browser runtime uses ES6 `import/export` syntax
        - to export from a module, include an export statement at the end of the module:
            - `export { functionA, functionB };`
            - you can export as default by creating an object which contains all export objects, then exporting the creatd object "as default"
            - e.g. `const rss = { funcA, funcB }; export {rss as default};`
        - import them using an import statement where needed:
            - `import { functionA, functionB } from './module.js';`
            - you can nickname imports to deal with name conflicts:
                - `inport { commonNameFunction as nickname } from './module.js';`
            - you can import defaults differently:
                - `import myImports from './module.js';`
                - `import { default as myImports } from './module.js';`
        - HTML scripts that import modules __MUST__ use the attribute `type="module"`

## Error Handling
- javascript errors are objects
    - have `name` and `message` properties
    - declared as `Error('error message')`
    - can use the `new` keyword as wellmistakes
- errors only stop the execution if they are thrown using the `throw` keyword
- use try-catch blocks for sections of code where we might anticipate an error

## Testing 
- manual testing
    - done by human users, often given a list of common actions to perform and expected behaviors
    - slow, costly, and prone to errors
### automated testing
- they use the `chai` library
- automated testing can double as documentation
- __regression__ is when the addition of a new feature breaks a previous feature (a.k.a. "the funcitonality regressed")
### types of testing
- from cheapest/fastest to expensive/slowest:
    - unit tests
        - covers smallest units (e.g. a function)
        - encompassing behavior/interaction/data should be *mocked*
    - integration tests
    - end-to-end tests
## Creating a node project
- install node and use npm to install mocha for testing
    - use npm init to create a project and generate a package.json file
    - after doing so, you can use npm install <package> -D to install and automatically add it as a dependency
### Node basics
- import/export syntax:
	- `const myModule = require('./modules/myModule.js');`
		- `functionA(var);`
	- `module.exports = { functionA, myArray };`
- Helpful libraries
	- `fs` (builtin) handles operations with file systems
		- `fs.appendFileSync(path, str)` creates a file at `path` with content `str` OR appends to file
		- `fs.readFileSync(path)` returns the contents of file at `path`
		- `fs.unlinkSync(path)` deletes the file found at `path`
	- `assert` (builtin) gives testing functions
		- `assert.equal(a, b);` evaluates 
	- `http` (builtin) handles HTTP servers and interactions
	- `ws` (not builtin) for creating websockets

### Creating a test suite (using node and Mocha)
- create a package.json file in the root directory, and include a json value like so:
    ` "scripts": {
        "test": "path/to/test/file.js"
    }, `
- then you can call npm test to run it
- if you are using mocha, you can set "test": "mocha" then use the npm test command ...
- you can use the `assert` library provided by Node for assert statements
    - `const assert = require('assert');`
    - assert.ok will throw an error if a passed argument evaluates to falsy
        - `assert.ok(4+4 === 8);`
    - assert.equals is a loose comparison
    	- assert.strictEqual is strict 
	- assert.deepStrictEqual compares the attributes of objects with strict equal
		- different objects with identical values are not strictly equal (objects, arrays, etc.)
- creating tests can be broken into 4 parts:
    - setup 
    - exercise
    - verify
    - teardown
### mocha
- uses "describe" and "it" statements
    - describe statements group tests (can be nested)
    - it statements explain a single test case
- both statements are functions which are called with two parameters:
    - a string to be used as the description
    - a callback function (usually anonymous) which contains the code to be executed for the test
        - tests are passed/failed using assert statements 
- take care not to mutate the testing environment in a test case
- use a "teardown" step at the end of each test case 
	- mocha provides hooks which can be used to handle similar teardowns
	- these hook functions are to be called inside the `describe` callback 
		- `beforeEach(callback)` runs before each test
		- `afterEach(callback)` runs after each test
		- before(callback)` runs before the first test
		- after(callback)` runs after the last test

## Test Driven Development
- the main cycle is: 
    - create test case first (it fails)
    - create the minimum possible code to make it pass (e.g. return expected value)
    - refactor the code after the test passes 
- the process is that you continually go trough the "red-green-refactor" cycle until you are confident in your coverage
    - dont forget edge cases

## Code Coverage & Test Coverage
### Code Coverage
- code coverage is the amount of the application that has been covered by test cases (as a %)
- depending on the organiation, they will include criteria such as:
    - has every function been called?
        - has every statement been executed?
            - has each boolean sub-expression evaluated as both true and false?
                - has every edge in the control flow graph been traversed?

### Test Coverage
- test coverage refers to the percentage of required features that have been tested (as opposed to the % of lines executed)
- e.g. writing tests for different devices and OS which should be supported

## Mocking
- mocking is creating a fake version of an external service for testing purposes
    - a.k.a 'stubbing', which is a slightly different process
- in integration tests, you should mock external services but not internal ones

## Spies & SinonJS
- spies are functions used to observe and record information about another function's calls
    - such as callCount, wasCalled, calledWith, etc.
- SinonJS is a library that includes standalone spies, fakes, and mocks
- there are a few ways to use spies, including function wrapping

## Async & HTTP
- code that is not asynchronous is considered "blocked" - meaning it must wait for other code to finish executing
- despite JavaScript being a single threaded language, it uses the "event loop" to simulate multithreading
- many callback functions are considered asynchronous because they only execute in response to other events
- `setTimeout(callback, ms)` can be used to delay the execution of the callback function by `ms` milliseconds
- `setInterval(callback, ms)` calls the function every `ms` milliseconds
- the `await` keyword will halt the execution of code (at the point of the word) and wait for a Promise to resolve
    - `await` can only be used inside an `async` function (but can be used freely in JS modules)
    - this allows the line directly after to use the result of the Promise
### The Event Loop
- the event loop is a way of emulating parallel execution using the Heap, Call Stack, web APIs, and an event queue
- web APIs handle asynchronous requests and pass things back to the stack using the event queue
- the events from the event queue are added to the stack periodically if the stack is empty

### Promises
- promises are objects which represent eventual outcomes
- promises have three states:
    - pending (initial)
    - fulfilled
    - rejected
- ex: 
    `const myPromise = new Promise((resolve, reject) => { });`
    - the resolve and reject functions can be called inside the promise's executor function to manually update the state of the promise and trigger the .then().catch() code defined elsewhere.
        - the argument passed with the resolve() and reject() is usable in the .then() and .catch() 
            - the .then() and .catch() must be called with anonymous functions defined to handle a parameter if you want to use one
- promises are consumed using .then() and .catch()
- .then() is a high order function which can take one, two, or zero callback functions as parameters
    - order is fulfilledHandler, rejectedHandler
    - .then() always returns a promise
- "composition": oftentimes, promises will be chained together. One promise resolution may invoke another promise etc.
    - when composing promise chains, you will likely be __returning__ a promise (or promise function rather) inside each .then() success handler. In this case, the next .then() success handler will be in response to the previously returned promise.
- common mistakes include:
    - nesting promises instead of chaining them
    - forgetting to return a promise (this is necessary for the next .then() to execute)








