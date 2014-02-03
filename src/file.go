package main

import (
	"io/ioutil"
)

type File struct {
	name string
}

func (f *File) SetName(name string) {
	f.name = name
}

func (f File) Name() string {
	return f.name
}

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