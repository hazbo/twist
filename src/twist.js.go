package main

func WrapJavascript() string {
	js := `

var editor = {
	sayHello : function(name) {
		sayHello(name);
	}
}

var e = editor;

	`
	return js
}