# Vue Notes
*see ./../exampleProjects/learning_vue/index.html*

## Contents
- [Vue Notes](#vue-notes)
  - [Contents](#contents)
  - [The Vue Instance](#the-vue-instance)
    - ["$ methods"](#-methods)
    - [Vue instance lifecycle hooks](#vue-instance-lifecycle-hooks)
  - [Templates](#templates)
  - [Components](#components)
  - [Directives](#directives)

## The Vue Instance
- every application needs a Vue instance, instantialized with:

            var vm = new Vue({
                //options object goes here
            })

- the main Vue component is typically organized into a tree of nested, reusable components
  - e.g.
    - Root Instance
      - TodoList
        - TodoItem
          - TodoButtonDelete
          - TodoButtonEdit
        - TodoListFooter
          - TodosButtonClear
          - TodoListStatistics
- all Vue components are also Vue instances, so they accept the same Options object
- when a Vue instance is created, all the properties in its `data` object are added to the reactivity system
  - when reactive values change, the view will update
  - this is only true for values that existed at first - values added to the Data object later on will not automatically trigger view updates
    - sometimes you may declare an empty value at first, expecting data later (e.g. `error: null`)
### "$ methods"
- Vue-specific instance properties and methods that are prefixed with `$` to differentiate them
- e.g. `vm.$data` returns the Data object used, `vm.$el` returns the root element used
- `vm.$watch('propertyOfVm', callbackfn(newVal, oldVal))` the watch method is used to set up a callback function to execute when a specified property of the Vue instance changes  
### Vue instance lifecycle hooks
- throughout the lifecycle of a Vue instance (initializations and whatnot) a series of hooks are available
- the "created" hook is run when the instance is created, and must be added to the Options object as `created: function () { ...`
- other hooks include `mounted`, `updated`, and `destroyed`
- note: do not use arrow syntax for these function definitions (it will make `this` no longer refer to the Vue instance)

![fucking brilliant](>./../../notesImages/vueinstancelifecycle.png)

## Templates
- use html syntax and allow for "mustache syntax" for interpolation: `{{ a }}` 
- mustache syntax also allows for javascript execution
  - e.g. `{{ num + 1 }}`

## Components

## Directives
- `v-once` (no parameter): will not update any normally reactive values inside the element
- `v-if="x < 5"`: will not add the element to the view if given condition is falsy
  - can be followed by an elt with `v-else`
- `v-html="<p>some raw HTML</p>"` allows you to insert raw HTML 
  - this raw HTML replaces the contents of the elt
    - be careful not to allow users to input HTML for this (XSS vulnerability)
- `v-bind:value="dynamicValue"` is used in HTML attributes instead of mustaches 
  - this is different for booleans, where false values will cause the attribute to not be rendered with the elt
- `v-for` TODO