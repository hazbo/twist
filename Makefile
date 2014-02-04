all: twist

bindata: bindatac
	tools/build-tools/go-bindata/bindata -o src/twist.js.go src/js

bindatac: tools/build-tools/go-bindata/main.go
	cd tools/build-tools/go-bindata/ && go build -o bindata

twist: src/twist.go bindata
	cd src && go build -o ../build/twist

clean:
	rm -f build/twist
	rm -f tools/build-tools/go-bindata/bindata
	rm -f build/*.txt
