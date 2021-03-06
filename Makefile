BINDATA_PATH=tools/build-tools/go-bindata/
BINDATA_SOURCE=tools/build-tools/go-bindata/main.go
BINDATA_BIN=tools/build-tools/go-bindata/bindata

INSTALL_PATH=/usr/local/bin

all: twist

bindata: bindatac
	${BINDATA_BIN} -o src/twist.js.go src/js/...

bindatac: ${BINDATA_SOURCE}
	cd ${BINDATA_PATH} && go build -o bindata

twist: src/twist.go bindata
	cd src && go build -o ../build/twist

install: build/twist
	cp build/twist ${INSTALL_PATH}

clean:
	rm -f build/twist
	rm -f tools/build-tools/go-bindata/bindata
	rm -f build/*.txt
