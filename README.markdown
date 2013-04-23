# gotags

## Tags

I got tired of exuberant tags failing on Ruby code, so I wrote a
version in go.

Features:

* Ruby only
* Emacs TAGS file only
* No command line options
* Fast

Detects:

* Classes and Modules
* Methods
* Constants
* attr_reader, attr_writer and attr_accessor definitions
* Aliases

## Usage:

```
$ gotags
```

or

```
$ gotags directories...
```
