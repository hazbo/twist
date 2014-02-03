all: twist

twist: src/goedit.go
	cd src && go build -o ../build/twist

clean:
	rm -f build/twist
	rm -f build/*.txt