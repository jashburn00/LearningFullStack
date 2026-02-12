# React

## The Virtual DOM
- browsers re-render the page for each detected change
  - it would be faster to only re-render the things that changed, *but* comparing DOMs is slow
- a virtual DOM is a JavaScript object that represents the tree-like structure of the actual DOM
  - comparing JavaScript objects is much faster than DOMs
- the Virtual DOM diffs are reported to the browser, and only necessary parts are changed
  - this prevents unnecessary re-rendering

## JSX
```
const h1 = <h1>Hello world</h1>;
```
![meme](2r3xxq.jpg)

- this is an example of __JSX__ 
- JSX code cannot be interpreted by normal browsers because it is an extension of JavaScript syntax
  - JSX code must be compiled to regular JavaScript by a JSX compiler before it reaches the browser
- JSX elements must have exactly one outermost element
- JSX elements are treated as javascript expressions, meaning they can "go anywhere javascript code can go"
- they can be given attributes just like HTML elements
  - `const p1 = <p id="large">foo</p>;`
- JSX elements can be nested, like normal HTML
  - if the JSX uses more than one line, it must be wrapped in `()`
- __Differences between JSX and HTML__:
  - JSX uses `className` instead of the `class` property
    - ex: `<h1 className="big red">Hello world</h1>`
  - self-closing tags in HTML are fine without the closing slash: `<br>`
    - JSX elements __must__ include the closing slash `<br/>`

## Rendering JSX
- It revolves around a __root__ object, which uses the method `render()` to display JSX elements
  - roots in React are created using `createRoot(elt)` which takes a DOM element as an argument
- the package `ReactDOM` is also needed for interacting with the DOM
    - typically done with: `import ReactDOM from 'react-dom/client';`
- You only need to call render() once. After the first call, React will manage updates to the DOM
    - typically, this first call is made from index.js and the rest of the updates are done in App.js

## using JavaScript inside JSX
- you must use curly braces
  - ex: `const mathproblem = <div>2 + 3 = {2+3}</div>;`
- JavaScript variables can be used in JSX expressions
  - note that javascript scope applies to these variables as it would any others
  - variables are commonly used to set attributes
    - ex:
    ```
    const [w, l] = [200, 650];
    const pics = {
        logo: "http://buckets.com/img/28",
        art: "https://buckets.com/img/10"
    }
    const logo = (
        <img
            src={pics.logo}
            alt="logo"
            height={l}
            width={w} 
        />
    );
    ```
## Event Listeners in JSX
- JSX elements have event listeners just like HTML
- in JSX, properties are given to assign event listening 
    - ex: `const img = <img onClick={clickHandler} src={path}/>;`
- example:
    - onClick 
    - onMouseOver
- JSX elements __cannot__ contain any `if` statements, but does handle ternary operator (a?b:c) and `&&`, `||`
    - you can conditionally add an element in JSX using embedded `&&`
        - ex: `{myBool && <div>element to render</div> }`
    - you can use `.map()` to add many elements to a list and then render them or add them to JSX elements
        - this works because JSX accepts lists of elements in some cases (such as ul/li)
        - sometimes lists of elements will require keys
            - `key` is an attribute in JSX elements, and it is used because sometimes React will scramble the order of lists
            - it is also used for elements that need to be "remembered" next time they are rendered
            - keys must be a unique value (shocker) 
- React components can be created without JSX using `React.createElement()`. This is secretly done for every JSX element 
    - createElement() takes parameters for type, ???, and value

## React Components
- React apps are built using components
    - components are small, reusable pieces of code that are responsible for one job (usually rendering some HTML and change it when some data changes)
    - example component:
    ```
    function myComponent(){
        return <h1>Hello world</h1>;
    }

    root.render(<myComponent/>);
    ```
- __function components__: use javascript functions to create React components
    - return a JSX element 
    - functional components usually follow PascalCase naming convention
- typically components are created in App.js and exported to index.js for use
- after creating, exporting, and importing a function component, you can simply use it as if it were a type of JSX element:
```
<MyComponent>poop</MyComponent>
//OR
root.render(<MyComponent/>);
```
- including external values in components can be done by using javascript
    - you must use variables reachable in the current scope, as normal
- functional components can be created using event handlers 
    - it is easily done by defining the event handler function before returning the component (inside the functional component function)
- __do not include `.js` at the end of the component path__

## Creating React Apps
- use the `npm` (or `npx` which comes with npm5.2 and newer) command `npx create-react-app MyApp` to create a boilerplate React app 
    - this comes complete with a .gitignore, package.json and package-lock, node modules
        - npm commands are set up initially as well
            - `npm start` to start the server
- *but how in the Aladeen does it work?*
  - it compiles using __webpack__ and serves the app files for access (localhost port 3000 by default)
## Using CSS with React
- surprisingly simple
- use a css file: `App.css`
- import it in the main app file: `import './App.css'; //inside App.js`

## Props & Component Interactions
- you can use other functional components inside functional component definitions (not that crazy) 
- props are used to pass information between components
  - requires a parameter in the receiving component
  - automatically sent in the form of key=value pairs when used as properties 
```
return <MyComponent name="dookie" />
// now the MyComponent(props) will have access to {props.name}
```
  - for non-string properties, you must use curly braces 
    - ex: `<MyComp age={98} stuff={['frog', 'dog']} />`
  - props values are immutable
- props is often used to pass information to child components
  - props are passed from parent to child to grandchild etc. in a top-to-bottom fashion
  - props are often used outside of display (e.g. login info, data)
- functions can be passed in with props
  - like we do with onClick, you can pass the name (not the signature) of functions into props 
  - commonly used for event handlers
    - naming conventions for passing event handlers in props:
      - the name of the property should be `onEvent` (where event is the event being handled)
      - the name of the handler function should be `handleEvent`
      - ex: `<MyComponent onJump={jumpHandler} />` 
      - be careful to not mix up your own props-based `onClick` with HTML's similarly named attributes
        - when using component instances, cases like this ^ result in props being created instead of  an HTML event listener 
- props always has a property called `children`, which returns all child elements of the component
    - either as an array or as a single element
- adding defaults to props can be done in three different ways:
    - specifying `MyComponent.defaultProps` property
        - uses object notation 
    - destructuring in parameter definition (like normal functions)
    - setting defaults in function body (like normal variables)
        - ex: `const {text = 'default text'} = props.text;`
## React Dev Tools
- the chrome extension 'React Developer Tools' can be used on web pages which use React
- the extension will display a colored logo if the page you're viewing uses React
    - blue logo means the page is a production deployment 
    - orange means development deployment
- the extension contributes two tabs to the chrome dev tools:
    - Components
        - this is similar to HTML elements in web inspector but for react components
        - you may need to clear some filters to see all components
    - Profiler
## Hooks
- hooks are React functions that can handle and manage states during runtime
- the main hooks are:
    - useState
        - named import from the React library so it uses destructuring import (e.g. `import React, { useState } from 'react';`)
        - returns an array with the initial state value and the state setter function 
            - e.g. `const [toggle, setToggle] = useState(); //accepts argument for initial state`
            - then you can use the setter to change the state which triggers a re-render
            - e.g. `setToggle(false);`
            - often used inside functional component definitions
        - ex: a functional component which returns a few buttons, each having onClick set to an arrow function changing the state of the color of the outder div's style 
        - state updates are __asynchronous__, and can potentially cause unexpected behavior (e.g. incrementing an integer twice asynchronously resulting in only incrementing by one)
            - this can be prevented by setting states using a callback function rather than explicit values
            - this makes use of state's previous value availability (ex: `setCount(curr => curr+1);`)
            - this is especially useful for non-static values or values that depend on previous state
    - useEffect
        - also a named export from the react library: `{ useEffect }`
        - mostly used for fetch data from backend, subscribe to data streams, manage timers/intervals, read/change the DOM
        - three main moments when useEffect is used:
            - when a component is initially rendered (aka mounted) to the DOM
            - when state or props change, causing a re-render
            - when the component is removed (aka unmounted) from the DOM
        - takes parameter for a callbackfn to execute in the three scenarios ^
        - expects a return value of a function to be executed on each re-render and upon unmounting
            - this function is used for teardown (e.g. for event listeners) to avoid memory leaks
        - example: (inside component function) `useEffect(somefunction); //where somefunction returns the cleanup function`
    - useContext
    - useReducer
    - useRef


