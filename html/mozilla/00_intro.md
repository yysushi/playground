# Note for HTML Introduction

## links

<https://developer.mozilla.org/ja/docs/Learn/HTML/Introduction_to_HTML>

## Notes

- html is a markup language
  - [markup](https://en.wikipedia.org/wiki/Markup_language) annotates a document in a way that is syntactically distinguishable from the text
- html indicates how the website should look like, to browser
- a html file consists of elements
  - element surrounds a content with start/end tags
  - e.g. `<p>hello world</p>`
    - the content "hello world" is surrounded with \<p\> element

- element can be nested with other element
  - e.g. `<p>hello <strong>world</strong></p>`
- two categories for elements
  - block-level element
  - inline-level element
  - e.g. "p" element is block-level element, so it inserts newline prior to next content.
  - this depends on CSS style

- empty element (sometimes, void element) can't be followed with end tag.
  - e.g. `<img src="https://raw.githubusercontent.com/mdn/beginner-html-site/gh-pages/images/firefox-icon.png">`

- element can have attribute which has two parts of attribute name and attribute value
  - e.g. `<a href="http://example.com">here</a>`
    - here is linked to somewhere with anker href attribute

- boolean attribute is an attribute without attribute value
  - e.g. `<input type="text" disabled>`

- html structure

```html
<!DOCTYPE html>
<html lang="en">
<head>
	<meta charset="UTF-8">
	<title></title>
</head>
<body>
</body>
</html>
```

- spaces in a row is recognized as one space
- there are special characters like `<>"'&`
  - it can be expessed `&lt;`, `&gt;`, `&quot;`, `&apos;` and `&amp;`
