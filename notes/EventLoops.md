# Event Loops

- primarily used to simulate asynchronous task delegation in single-threaded execution
- typically:
    - execute javascript on a single main thread
    - use a call stack
    - maintain a task queue (macrotasks)
    - maintain a microtask queue
    - defer I/O or timer work to host APIs (web APIs, libuv)
    - drain microtasks before moving to next macrotask
- basic concept:
    - process all microtasks
    - take next macrotask
    - repeat

## JavaScript/Browser Event Loop
- basic loop:
    - execute one macrotask (script, timer callback, event handler, etc.)
    - drain the microtask queue completely
    - render (if needed)
    - repeat
- macrotasks include:
    - `setTimeout` and `setInterval`
    - DOM events
    - `postMessage`
    - network callbacks
    - script execution
- microtasks:
  - `Promise.then()`
  - `queueMicrotask`
  - `MutationObserver`
- Browser specific behavior:
  - rendering happens between macrotasks (if necessary)
  - microtasks are completed before rendering 
  - no "phase model" available to the developer

## Node Event Loop
- main differences:
  - Node event loop is implemented by __libuv__ 
  - it has explicit phases
- each loop iteration consists of ordered phases:
  1. __timers__ (`setTimeout`, `setInterval`) 
  2. __pending callbacks__ (system-level I/O callbacks)
  3. __idle, prepare__
  4. __poll__ (retrieve new I/O events, execute I/O callbacks, stop here if nothing else scheduled)
  5. __check__ (`setImmediate` callbacks)
  6. __close callbacks__ (e.g. `socket.on('close')`)
- differences in microtask handling:
  - microtasks are processed between each phase
  - Node specifically has two microtask queues:
    - Promise microtask queue
    - `process.nextTick()` queue (higher priority)
      - this is a Node-specific queue and is not present in browsers
      - runs immediately after current operation