# HTML Notes

## Contents
<!-- no toc -->
- [Main Structure of HTML Pages](#mainstructure)
- [Mozilla Developer Network](#mdn)
- [Debugging](#debugging)
- [Tables](#tables)
- [Semantics](#semantics)

## <a id="mainstructure">Main Structure of HTML Pages</a>
- __mandatory prefix__ < !DOCTYPE html> lets the browser know it needs to render an HTML document as well as the version of HTML ("html" defaults to current standard)
- the entire rest of the file needs to be __enclosed in html tags__ e.g. < html> </ html>
  - __header__ 
    - __contains metadata__ (such as a stylesheet reference or page title) that isn't displayed on the page
  - __body__
    - content must be rendered inside the body tag (though some browsers still render elements outside it, it is considered bad practice)
    - typically used for the visible content of a webpage

## <a id="mdn">Mozilla Developer Network</a>
site: [mozilla docs](https://developer.mozilla.org/en-US/)

## <a id="debugging">Debugging</a>
- W3C has a [Markup Validation Service](https://validator.w3.org/) to help debug HTML code

## <a id="tables">Tables</a>
table structure:
- table
  - thead
    - tr
      - th (for each header/column)
  - tbody
    - tr
      - th or td (for each column)
<!-- end of table structure --> 
- each th can be given an attribute *scope*, set to either row or col
- td or tr that spans several cols can be given the *colspan* attribute with value >= 1
- td or tr that spans several rows can be given the *rowspan* attribute with value >= 1

## <a id="semantics">Semantics</a>
- semantic HTML elements provide context information regarding the elements they encapsulate
  - e.g. "div" is not semantic, "header" is semantic
- benefits:
  - search engine optimization
  - accessibility
  - readability

## Layouts (Grid)
- grid layout is comprised of columns and rows evenly spanning across the page
  - commonly 12, 16, or 3 cols
  - rows can be used to force content away from empty space 
  - responsive grids are needed to handle the 
- gutters are the negative space between columns
  - no content is allowed in gutters unless it is using both adjacent columns
  - gutters should be much smaller than columns or rows
- margins on either side of the design (leftmost and rightmost cols?)
- whitespace
  - passive/micro whitespace: improves aesthetics without guiding users
  - active/macro whitespace: done intentionally to emphasize content or guide users