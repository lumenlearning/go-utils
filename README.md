go-utils
========

'go-utils' is a set of utility functions for the Go programming language that
facilitate the work we are trying to accomplish at Lumen.

This repo contains the following Go packages:

- *canvas*: A number of functions that are useful for working with both the
API and web interfaces for the Canvas Learning Management System from
Instructure.
- *html*: These functions will help with finding specific tags in an HTML document and extracting the text from those tags.
- *http*: These package helps with common HTTP functions and includes a
UserAgentClient struct that allows you to make HTTP requests using a specified
User-Agent header.
- *io*: This package currently contains only one convenience function: ReadFile.
- *time*: This package contains several utility functions for converting
between two common time formats and is helpful for working with page view data
from Canvas.
