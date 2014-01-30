package main

type Line struct {
	count int
}

func (l *Line) addLine() {
	l.count++
}

func (l *Line) deleteLine() {
	l.count--
}