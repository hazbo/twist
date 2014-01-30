package main

type EditorBuffer struct {
	content []string
}

type Line struct {
	count int
}

func (b *EditorBuffer) appendToStackBuffer(char []string) {
    c := make([]string, len(b.content) + len(char)) 
    copy(c[:len(b.content)], b.content) 
    copy(c[len(b.content):], char) 
    b.content = c
}

func (b *EditorBuffer) detachFromStackBuffer() {
	content_len := len(b.content)
	if (content_len > 0) {
		b.content = b.content[:content_len-1]
	}
}

func (l *Line) addLine() {
	l.count++
}

func (l *Line) deleteLine() {
	l.count--
}