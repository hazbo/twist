# Twist

### Note
This is all just an experiment. There are very basic things about this editor
such as simply deleting a line, which has not been implemented! It's more just
a concept of what I'd maybe one day like to build fully.

### About
Twist is a terminal text editor written in [Go](http://golang.org/). I am fairly new to Go
and so far this project has been a nice learning curve. There may be some things that suck
about it; so if you feel you can contribute to improve any aspect of it, please do.

One of the things I wanted to add to this editor is a Javascript API that will
allow you to perform various tasks that could / should help you save time.

I decided I'd have a shot at making a simple terminal editor after watching
[A Whole New World](https://www.destroyallsoftware.com/talks/a-whole-new-world) by
Gary Bernhardt. It doesn't actually implement a lot of what he is talking about
in it at all, but was inspirational.

### Install

GoEdit depends on [gocui](https://github.com/jroimartin/gocui) which in turn depends on
[termbox-go](https://github.com/nsf/termbox-go). These are actually forks stored locally
in the `vendor` directory so they are easier to maintain and customise; hence the `submodule init`.

The forked versions along with other forked dependencies can be found here:

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

### PHP Support

Because I use PHP on a daily basis I decided to see if I could start adding things to Twist
that would maybe be relevant to some of the PHP work I do. As of now, in terms of PHP
support, it does next to nothing. I'm not the biggest fan of text editors / IDEs genenrating
loads of code for you behind the scenes or always trying to guess what you're going to do next
and then getting it wrong. But I did want a little bit of code genenration in here not only that
but for the editor to be aware of this code and to know what has happened so it can become more useful
as you keep using it.

An example:

There is a [horrible bit of JavaScript](https://github.com/hazbo/twist/blob/master/src/js/lang/php.js#L33-L46)
that will generate you a new, empty PHP class (given that there is nothing currently in the
editor at this point). But due to:

```javascript
this.class     = className;
this.namespace = namespace;
```

Twist now knows what your class is and the namespace for that class. Not that it can do
anything with that information, yet...

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

As mentioned earlier, I did try to add a little PHP support, just to see if it
could be done nicely. The following line would generate you a new PHP class:

```javascript
e.print(php.newClass('MyNamespace', 'MyNewClassName'));
```

And then Twist is now aware of this namespace and class, so we can use that
for something else later on perhaps in Javascript. An easy way to test this
is to simply just print it out:

```javascript
e.print(php.class);
```

### Basic Usage

#### Commands

  - New file: `Ctrl+N`
  - Quit: `Ctrl+C`
  - Editor mode: `Ctrl+K`
  - JS mode: `Ctrl+J`
  - Save (write) a file: `Ctrl+W` or `e.write('filename.txt');` in the console

#### Javascript tricks

`e` is always just an alias for `editor`. Either will work:

  -  `e.print("Print some text to the buffer");`
  -  `e.println("Print a new line to the buffer");`
  -  `e.write("Newfile.txt");`
  -  `e.print(php.newClass("ClassNameSpace", "ClassName"));`
  -  `e.print(php.newFunction("FunctionName"));`

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