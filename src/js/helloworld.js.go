package main

import (
	"../../vendor/otto"
)

func ExecuteGoJs() {
	otto := otto.New()
	otto.Run(`

function SayHello()
{
	console.log('Hello, World!')
}

	`)
}