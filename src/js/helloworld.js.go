package main

import (
	"../../vendor/otto"
)

func execute() {
	otto := otto.New()
	otto.Run(`

function SayHello()
{
	console.log('Hello, World!')
}

	`)
}