# HTML Notes

## Contents
<!-- no toc -->
- [Main Structure of HTML Pages](#mainstructure)
- [Browser Events](#browser-events)
- [List of Common Elements](#commonelements)

## <a id="mainstructure">Main Structure of HTML Pages</a>
- __header__ 
- __body__
  - content must be rendered inside the body tag (though some browsers still render elements outside it, it is considered bad practice)

## <a id="browser-events">Browser Events</a>
- onmouseover (mouse near an HTML element)

## <a id="commonelements">List of Common Elements</a>
- __headings (h1 - h6)__ 1 is biggest
- __div (div)__ used to containerize content
  - uses the *id* attribute for dynamic lookup
    - all elements may use this attribute
- __span (span)__ used to specify inline blocks of text
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