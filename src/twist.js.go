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
		0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x00, 0xff, 0x2a, 0x4b,
		0x2c, 0x52, 0x48, 0x4d, 0xc9, 0x2c, 0xc9, 0x2f, 0x52, 0xb0, 0x55, 0xa8,
		0xe6, 0xe2, 0xe2, 0x2c, 0x28, 0xca, 0xcc, 0x2b, 0x51, 0xb0, 0x52, 0x48,
		0x2b, 0xcd, 0x4b, 0x2e, 0xc9, 0xcc, 0xcf, 0xd3, 0x28, 0x49, 0xad, 0x28,
		0xd1, 0x04, 0x4a, 0x71, 0x42, 0xa4, 0x20, 0x7c, 0x2e, 0xce, 0x5a, 0x98,
		0xda, 0x9c, 0x3c, 0x42, 0xaa, 0x15, 0xb4, 0x15, 0x94, 0x62, 0xf2, 0x94,
		0x20, 0x9a, 0x80, 0x08, 0x6c, 0x27, 0xd0, 0x3a, 0x88, 0xbd, 0xd6, 0x80,
		0x00, 0x00, 0x00, 0xff, 0xff, 0x7e, 0xe4, 0x79, 0xd7, 0x83, 0x00, 0x00,
		0x00,
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
