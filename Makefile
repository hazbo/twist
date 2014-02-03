all: twist

twist: src/twist.go
	cd src && go build -o ../build/twist

clean:
	rm -f build/twist
	rm -f build/*.txt