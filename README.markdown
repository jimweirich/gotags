# gotags

## Tags

I got tired of exuberant ctags failing on Ruby code, so I wrote a
version in go.

Features:

* Ruby only
* Emacs TAGS file only
* No command line options
* But it's **fast**

Detects:

* Classes and Modules
* Methods
* Constants
* attr_reader, attr_writer and attr_accessor definitions
* Aliases

## Usage:

```
$ gotags file [file...]
```

Analyze all files listed in the command line. If the file is a
directory, then recursively analyze all the files in that directory.
If no files are listed, then the current directory ('.') is assumed.

Only <code>.rb</code> files are actually handled, all other files are
silently ignored.
