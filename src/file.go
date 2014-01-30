package main

import (
	"io/ioutil"
)

func GetFileContents(filename string) []byte {
	b, err := ioutil.ReadFile(filename)
    if err != nil {
            panic(err)
    }
    return b
}