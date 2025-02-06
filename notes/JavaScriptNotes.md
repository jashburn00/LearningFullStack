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







