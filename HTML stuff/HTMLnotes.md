# HTML Notes

## Contents
<!-- no toc -->
- [Main Structure of HTML Pages](#mainstructure)
- [High Level and Overview Info](#highlevel)
- [Browser Events](#browser-events)
- [List of Common Elements](#commonelements)

## <a id="mainstructure">Main Structure of HTML Pages</a>
- __mandatory prefix__ < !DOCTYPE html> lets the browser know it needs to render an HTML document as well as the version of HTML ("html" defaults to current standard)
- the entire rest of the file needs to be __enclosed in html tags__ e.g. < html> </ html>
  - __header__ 
    - __contains metadata__ (such as a stylesheet reference or page title) that isn't displayed on the page
  - __body__
    - content must be rendered inside the body tag (though some browsers still render elements outside it, it is considered bad practice)
    - typically used for the visible content of a webpage

## <a id="highlevel">High Level & Overview Info</a>
- references between pages are assuming of the typical root directory structure

## <a id="browser-events">Browser Events</a>
- onmouseover (mouse near an HTML element)

## <a id="commonelements">List of Common Elements</a>
- __headings (h1 - h6)__ 1 is biggest
- __div (div)__ used to containerize content
  - uses the *id* attribute for dynamic lookup
    - all elements may use this attribute
- __span (span)__ used to specify inline blocks of text
- __break (br)__ used to add a line break 
- __emphasized text (em)__ italics
- __strong text (strong)__ bold
- __unordered lists (ul)__ 
  - contain __list items (li)__ 
- __ordered lists (ol)__
  - also contain __list items (li)__
- __image (img)__ 
  - uses a self-closing tag, meaning the final forward slash may be omitted: < img /> or < img > 
  - uses the *src* attribute to specify a URL
  - uses the *alt* attribute to specify text displayed when the image is unable to load, to be used as an accessibility feature, and for search engine optimization
    - this attribute should be left empty if the image is not significant (e.g. a logo)
- __video (video)__ uses *src*,  *width*, *height*, and *controls* attributes
  - requires closing tag unlike images
  - text between tags is displayed if video is unable to load
  - ex: < video src="mymovie.mp4" width="240" height="160" controls> Video not found. < /video >
- __anchors/links (a)__ uses *href* attribute and uses display text
  - uses the *target* attribute to specify opening in new tab/window (_blank) or ...
  - other elements can be wrapped in an anchor element to use photo links, etc.
    - in the case of __li__, only wrap the innermost text
  - can target an element on the same page by using the *id* property of the target element, prefixed with a "#" (e.g. href="#mydiv")