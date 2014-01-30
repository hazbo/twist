package main

import (
	"os"
)

func GetFilenameArg() string {
	return os.Args[1]
}