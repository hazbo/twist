package main

type Buffer struct {
	content string
}

type Line struct {
	count int
}

func (l *Line) addLine() {
	l.count++
}

func (l *Line) deleteLine() {
	l.count--
}