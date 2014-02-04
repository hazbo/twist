package main

import (
    "bytes"
    "compress/gzip"
    "fmt"
    "io"
)

func bindata_read(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	return buf.Bytes(), nil
}

func src_js_twist_js() ([]byte, error) {
	return bindata_read([]byte{
		0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x00, 0xff, 0x4a, 0x2b,
		0xcd, 0x4b, 0x2e, 0xc9, 0xcc, 0xcf, 0x53, 0x48, 0x2f, 0x4a, 0x4d, 0x2d,
		0xd1, 0xd0, 0x54, 0xa8, 0xe6, 0xe2, 0x2c, 0x4a, 0x2d, 0x29, 0x2d, 0xca,
		0x53, 0x50, 0xf2, 0x48, 0xcd, 0xc9, 0xc9, 0xd7, 0x51, 0x08, 0xcf, 0x2f,
		0xca, 0x49, 0x51, 0x54, 0xb2, 0xe6, 0xaa, 0x05, 0x04, 0x00, 0x00, 0xff,
		0xff, 0x21, 0x89, 0xcc, 0x63, 0x2d, 0x00, 0x00, 0x00,
		},
		"src/js/twist.js",
	)
}


// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	if f, ok := _bindata[name]; ok {
		return f()
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string] func() ([]byte, error) {
	"src/js/twist.js": src_js_twist_js,

}
