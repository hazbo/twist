# Twist

### About
Twist is a terminal text editor written in [Go](http://golang.org/). I am fairly new to Go
and so far this project has been a nice learning curve. There may be some things that suck
about it; so if you feel you can contribute to improve any aspect of it, please do.

One of the things I wanted to add to this editor is a Javascript API that will
allow you to perform various tasks that could / should help you save time.

### Install

GoEdit depends on [gocui](https://github.com/jroimartin/gocui) which in turn depends on
[termbox-go](https://github.com/nsf/termbox-go). These are actually forks stored locally
in the `vendor` directory so they are easier to maintain and customise; hence the `submodule init`.

The forked versions along with other dependencies can be found here:

  - [gocui](https://github.com/hazbo/gocui)
  - [termbox-go](https://github.com/hazbo/termbox-go)
  - [otto](https://github.com/hazbo/otto.git)
  - [go-bindata](https://github.com/hazbo/go-bindata)

To build from source and get everything going all you need are the commands below. These will
sort out the dependencies for you.

	$ git clone https://github.com/hazbo/twist.git
	$ cd twist
	$ git submodule init
	$ git submodule update
	$ tools/configure
	$ make
	$ build/twist

![v0.0.1](https://raw.github.com/hazbo/twist/master/screenshots/v0.1/3.png?token=315774__eyJzY29wZSI6IlJhd0Jsb2I6aGF6Ym8vdHdpc3QvbWFzdGVyL3NjcmVlbnNob3RzL3YwLjEvMy5wbmciLCJleHBpcmVzIjoxMzkyMTI3OTEzfQ%3D%3D--624c90f8058903d2e94fb7cf9697457120cd23f5)

### Javascript API

There are Javascript functions within the `console.go` written in Go with the help of
[otto](https://github.com/hazbo/otto.git). On top of those is what will be an API that
allows the editor to interact with some of the things performed using Go. This is a
standard Javascript file found in `src/js/twist.js`. Anything that is put in here will
be executed at run time as you start using the Javascript console.

There is virtually next to nothing in there at the moment. Things will be added to
allow you to perform all sorts of tasks but here is a quick example anyway:

```javascript
jsc >> e.print("Hello, World!");
```

This will just append the string "Hello, World" to your buffer. Another example
using a loop could look like this:

```javascript
jsc >> for (var i = 0; i < 10; i++) { e.println('Hello, World!'); }
```

And as you would expect it will be append to your buffer multiple times.

I will come back to documenting the Javascript API when it is more useful.

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