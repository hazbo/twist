# Twist

### About
Twist is a terminal text editor written in [Go](http://golang.org/). I am fairly new to Go
and so far this project has been a nice learning curve. There may be some things that suck
about it; so if you feel you can contribute to improve any aspect of it, please do.

### Install

GoEdit depends on [gocui](https://github.com/jroimartin/gocui) which in turn depends on
[termbox-go](https://github.com/nsf/termbox-go). These are actually forks stored locally
in the `vendor` directory so they are easier to maintain and customise; hence the `submodule init`.

The forked versions along with other dependencies can be found here:

  - [gocui](https://github.com/hazbo/gocui)
  - [termbox-go](https://github.com/hazbo/termbox-go)
  - [otto](https://github.com/hazbo/otto.git)

To build from source and get everything going all you need are the commands below. These will
sort out the dependencies for you.

	$ git clone https://github.com/hazbo/twist.git
	$ cd twist
	$ git submodule init
	$ git submodule update
	$ tools/configure
	$ make
	$ build/twist

![v0.0.1](https://raw.github.com/hazbo/twist/master/screenshots/v0.0.1/1.png?token=315774__eyJzY29wZSI6IlJhd0Jsb2I6aGF6Ym8vdHdpc3QvbWFzdGVyL3NjcmVlbnNob3RzL3YwLjAuMS8xLnBuZyIsImV4cGlyZXMiOjEzOTIwNDI3NjF9--29f0d28d54a4c3bc0ce06c06619ae9392ea9eeac)

### Development Notes

There are various small shell scripts in the `tools` directory. As of now these are just in place
to keep dependencies up to date and ensuring that new changes to those dependencies will not
break things in the editor.

If you plan on contributing it may be a good idea to run:

	$ tools/upstream
	$ tools/pull

This will add the upstream repos for where the dependencies originate. Ideally commits made on these
repos will already be in our forked versions, but these tools can assist with that.

### Contributors:

  - [@harry4_](http://twitter.com/harry4_)

### Thanks To:

  - [Roi Martin](https://github.com/jroimartin) for [gocui](https://github.com/jroimartin/gocui)
  - [nsf](https://github.com/nsf) for [termbox-go](https://github.com/nsf/termbox-go)
  - [Robert Krimen](https://github.com/robertkrimen) for [otto](https://github.com/robertkrimen/otto)

### Contributing

  - Fork Twist
  - Create a feature branch (`git checkout -b my-feature`)
  - Commit your changes (`git commit -am 'Add some feature'`)
  - Push to your feature branch (`git push origin my-feature`)
  - Create new Pull Request