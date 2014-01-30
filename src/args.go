package main

import (
	"os"
)

func FilenameIsSet() bool {
	if len(os.Args) < 2 {
		return false
	}
	return true
}

func GetFilenameArg() string {
	return os.Args[1]
}