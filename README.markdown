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

Only <code>.rb</code>, <code>.rake</code> and <code>Rakefile</code>
files are actually handled, all other files are silently ignored.

## Building

Make sure you have a recent installation of the Go language on your
system.  Then do the following:

```
$ git checkout git://github.com/jimweirich/gotags.git
$ cd gotags
$ go install onestepback.org/gotags
$ cp bin/gotags SOMEWHERE_IN_YOUR_PATH
```

## License

Copyright 2013 by Jim Weirich

This software is available under the MIT license.  See MIT-LICENSE in
the repository for details.

## Links

* Source: http://github.com/jimweirich/gotags
* Clone URL: git://github.com/jimweirich/gotags.git
* Issues: https://github.com/jimweirich/gotags/issues
