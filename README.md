# GoEdit
This project needs a new name

### About
"GoEdit" is a terminal text editor written in Go.

### Install

GoEdit depends on [gocui](https://github.com/jroimartin/gocui) which in turn depends on
[termbox-go](https://github.com/nsf/termbox-go). These are actually forks stored locally
in the `vendor` directory so they are easier to maintain and customise; hence the `submodule init`.

The forked versions can be found here:

  - [gocui](https://github.com/hazbo/gocui)
  - [termbox-go](https://github.com/hazbo/termbox-go)

To build from source and get everything going all you need are the commands below. These will
sort out the dependencies for you.

	$ git clone https://github.com/hazbo/goedit.git
	$ cd goedit
	$ git submodule init
	$ git submodule update
	$ tools/configure
	$ make
	$ build/goedit

![v0.0.1](https://raw.github.com/hazbo/goedit/otto-integration/screenshots/v0.0.1/2.png?token=315774__eyJzY29wZSI6IlJhd0Jsb2I6aGF6Ym8vZ29lZGl0L290dG8taW50ZWdyYXRpb24vc2NyZWVuc2hvdHMvdjAuMC4xLzIucG5nIiwiZXhwaXJlcyI6MTM5MTcwNTAyMH0%3D--2a0493fc7c392f60d489e73864219d20d11b6e20)

### Development Notes

There are various small shell scripts in the `tools` directory. As of now these are just in place
to keep dependencies up to date and ensuring that new changes to those dependencies will not
break things in the editor.

If you plan on contributing it may be a good idea to run:

	$ tools/upstream
	$ tools/pull

This will add the upstream repos for where the dependencies originate. Ideally commits made on these
repos will already be in our forked versions, but these tools can assist with that.

### Contributors

  - [@harry4_](http://twitter.com/harry4_)

### Thanks To:

  - [Roi Martin](https://github.com/jroimartin) for [gocui](https://github.com/jroimartin/gocui)
  - [nsf](https://github.com/nsf) for [termbox-go](https://github.com/nsf/termbox-go)
  - [Robert Krimen](https://github.com/robertkrimen) for [otto](https://github.com/robertkrimen/otto)

### Contributing

  - Fork GoEdit
  - Create a feature branch (`git checkout -b my-feature`)
  - Commit your changes (`git commit -am 'Add some feature'`)
  - Push to your feature branch (`git push origin my-feature`)
  - Create new Pull Request