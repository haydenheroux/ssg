# ssg

Simple static page generator built using vanilla Go and written to be simple to understand.

## Usage

Build the binary and run on a single file, producing `input_file.html`.

```
go build
# Add the ssg binary to PATH
cd path/to/files
touch header_template footer_template
ssg input_file
```

Optionally, use the `-o` argument to manually specify the output file.

```
cd path/to/files
touch header_template footer_template
ssg input_file -o output_file
```

## Features

### Syntax

Source files for pages are written in plain text, and store only two types of information: text content, and element metadata. Text content is specified as text, and unless otherwise stated is always present in the document in some way; on the other hand, element metadata is only present in the source file, and usually indicates an element which contains more than text. For instance, compare a paragraph with an image.

```
Here is some text which will be included as a paragraph in the generated page.
```

```
@image path/to/image Here is a caption for the image.
```

While the paragraph simply contains the text of the paragraph, the image includes the metadata for which image to include. Generally, any line beginning with an "@" symbol contains some form of metadata, which may be followed by some amount of text content.

Paragraph.
```
# Paragraph
Here is a paragraph.
```

Code, copies contents of a file into the HTML `<code>` element, including line numbers within in a `<samp>` element.
```
# Code
@code path/to/code
@include path/to/code
```

Image, inserts HTML `<figure>` containing a sourced image, with the caption written within a `<figcaption>` element. 
```
# Image
@img path/to/image Here is a caption for the image.
@image path/to/image Here is a caption for the image.
```

Header, `<h2>`.
```
# Header
@section Here is a header for a section of the page.
@segment Here is a header for a segment of the page.
```

List, inserts a `<ul>` containing the line for each line in the specified file.
```
# List
@list path/to/list
```

### Formatting Rules

Pages generated are formatted using an opinionated formatter which enforces the following rules:

1. No lines with no content.
