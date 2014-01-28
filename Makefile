all: goedit

goedit: src/goedit.go
	cd src && go build -o ../build/goedit

clean:
	rm build/goedit