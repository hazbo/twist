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

func WriteFile(filename string, b []byte) {
	err := ioutil.WriteFile(filename, b, 0644)
	if err != nil{
		panic(err)
	}
}