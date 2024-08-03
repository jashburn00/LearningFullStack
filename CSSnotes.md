# CSS
## Contents
<!-- no toc -->
- [How to apply styles](#how-to-apply-styles)

## How to apply styles
- two main ways to apply styles to an element:
  - selectors
    - applied as a ruleset 
    - there are different types of selectors
      - type selector targets each element of one or more types (the universal selector, *, targets all elements on the page)
        > header { ...
      - class selection will apply to all elements given the selected class value (this one uses the '.' prefix)
        > .brand { ...
  - inline
    - applied to individual element
      > < p style="color: blue;">
    - each style should be punctuated ( ; ), even if there is only one
- both use *declarations*, which are the property name and value
- the brackets following the selector delimit the *declaration block*.
- 