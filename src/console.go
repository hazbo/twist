package main

import (
	"../vendor/otto"
)

func execute(javascript string) {
	Otto := otto.New()

	Otto.Run(`
		console.log('Hello World');
	`)
}