# Note for HTML header

## Links

<https://developer.mozilla.org/ja/docs/Learn/HTML/Introduction_to_HTML/The_head_metadata_in_HTML>

## Notes

- \<head\> element includes meta data
  - meta data is not visible in the page
  - one motivation of this page is to be familiar with some famous head elements

- \<title\> element is a page title in browser tab or bookmark tool bar
  - on the other hand, \<h1\> is a title markup in the browser page

- \<meta\> element adds meta data to the document
  - character encoding e.g. `<meta charset="utf-8">`
  - author and description with using name and content attribute e.g. `<meta name="author" content="Chris Mills">`
    - descritpion is used by search engine like google
  - link image by sns

- favorite icons e.g. `<link rel="shortcut icon" href="favicon.ico" type="image/x-icon">`

- add css and javascript
  - `<link rel="stylesheet" href="my-css-file.css">`
  - `<script src="my-js-file.js"></script>`

- \<link\> element is placed in head
- \<script\> element not needs to be placed in head, it is placed in the tail of body 
