# gotags

## Tags

I got tired of exuberant ctags failing on Ruby code (see
https://gist.github.com/jimweirich/5440424), so I wrote a version in
go.

Features:

* Ruby & Go only
* Emacs TAGS file only
* Minimal command line arguments
* But it's **fast**

Detects:

* Ruby (extensions <code>.rb</code> and <code>.rake</code>, also files named <code>Rakefile</code>)
  * Classes and Modules
  * Methods
  * Constants
  * attr_reader, attr_writer and attr_accessor definitions
  * Aliases

* Go (extension <code>.go</code>)
  * Top level functions and variables
  * Interfaces
  * Structs
  * Methods

## Usage:

```
$ gotags [options] file [file...]
```

Analyze all files listed in the command line. If the file is a
directory, then recursively analyze all the files in that directory.
If no files are listed, then the current directory ('.') is assumed.

Only the files and extensions listed above are actually handled, all
other files are silently ignored.

**Command Line Options:**

* <code>-v</code> -- Display the version of the program and exit.
* <code>-h</code> -- Display a help message.

## Building

Make sure you have a recent installation of the Go language on your
system.  Then do the following:

```
$ git clone git://github.com/jimweirich/gotags.git
$ cd gotags
$ go install onestepback.org/gotags
$ cp bin/gotags SOMEWHERE_IN_YOUR_PATH
```

See the [Links Section](#links) for links to binary executables.

## License

Copyright 2013 by Jim Weirich

This software is available under the MIT license.  See MIT-LICENSE in
the repository for details.

## Links

* Source: http://github.com/jimweirich/gotags
* Clone URL: git://github.com/jimweirich/gotags.git
* Issues: https://github.com/jimweirich/gotags/issues
* Platform executables:
  * [Mac OSX Version 1.1.3](http://onestepback.org/download/gotags-1.1.3-darwin-x86_64.tgz "TGZ File")
  * [Linux-x86 Version 1.1.1](http://onestepback.org/download/gotags-1.1.1-linux-x86_64.tgz "TGZ File")
