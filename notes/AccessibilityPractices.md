# Accessibility Practices (A11y)

Web accessibility is not regulated but has an evangelist body Web Content Accessibility Guidelines 2.1 (WCAG 2.1) that creates de facto standards.

Web accessibility is the practice of making web content usable and consumable by individuals with impairments or disabilities that would hinder their ability to use the web.

## Key Accessibility Practices

### Contrast
Constrast (specifically of color) is used to make all fonts readable and more easily so by impaired individuals.

### Headings
Headings are used to help screen readers and those who use them glean the overall structure of the page. The general rule of thumb is larger heading = more important. It is considered best practice to not skip heading numbers where possible.

### Text
- use real text instead of images of text
- text should be readable
- text color must be sufficiently contrasted with its background color
- text size should be 16px or greater

### Tag Usage and Types
- using the appropriate tag type can sometimes increase screen reader efficacy.
    - e.g. placing all header-related elements inside a < header> tag, not a < div class="header">

### ARIA
- not sure exactly what that is
- assign the 'role' attribute of elements to assist screen readers
    - e.g. role="complementary" for supporting information 
- some roles include:
    - complementary: supporting information to the main information
    - presentation: not intended to be read by screen readers (vital chidren are skipped as well, but optional children are still read)
- there's a separate property called 'aria-label' used to give screen readers additional info (intended to be read to user)

### alt attribute
- some elements have built-in 'alt' functionality
- the value assigned to the alt attribute will display if the element fails to load.
- this is typically used to provide backup annotations for images 
- in some cases alt should be empty, but it should never be omitted

