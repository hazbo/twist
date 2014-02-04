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
		0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x00, 0xff, 0x8c, 0x8f,
		0xb1, 0x4e, 0xc3, 0x30, 0x10, 0x86, 0x67, 0xfc, 0x14, 0xbf, 0x3a, 0xb5,
		0x01, 0x5a, 0x58, 0x8b, 0x98, 0x50, 0x11, 0x95, 0x02, 0x0b, 0x1d, 0x59,
		0x8c, 0x7d, 0x21, 0x27, 0xb9, 0x4e, 0xe4, 0x5c, 0x62, 0x2a, 0xd4, 0x77,
		0xc7, 0x4e, 0x0b, 0x94, 0xad, 0x27, 0x65, 0xc8, 0x77, 0xf7, 0xdd, 0xfd,
		0x56, 0x8b, 0xa2, 0x50, 0x28, 0xb0, 0x89, 0xdc, 0x09, 0x86, 0x9b, 0xf9,
		0xed, 0xb5, 0xa5, 0x21, 0x91, 0x0c, 0xa7, 0x66, 0x86, 0x27, 0x1d, 0xc2,
		0x0e, 0xa5, 0x8e, 0x81, 0xbc, 0xa1, 0x63, 0xa3, 0x64, 0x43, 0xbe, 0xa3,
		0x25, 0x9e, 0xd7, 0x9b, 0xfc, 0x9f, 0xbf, 0xc7, 0x26, 0x40, 0x6a, 0x42,
		0xd5, 0x3b, 0x07, 0xd3, 0xb4, 0xbb, 0xc0, 0x1f, 0xb5, 0x40, 0x7b, 0x0b,
		0x77, 0x18, 0x07, 0xfb, 0xaa, 0x09, 0x5b, 0x2d, 0xdc, 0xf8, 0x2b, 0xb4,
		0x8e, 0x74, 0x62, 0x03, 0x53, 0x1c, 0xbd, 0x72, 0xfd, 0xb0, 0x7a, 0x79,
		0x5d, 0xe5, 0x55, 0x15, 0x3b, 0x4a, 0x4c, 0x0b, 0xa2, 0xee, 0x60, 0x53,
		0xb2, 0xc0, 0xef, 0xbd, 0x90, 0x45, 0x64, 0xa9, 0x53, 0x87, 0x3b, 0x74,
		0x4d, 0x1f, 0x0c, 0xa5, 0x43, 0x96, 0xe6, 0xc9, 0x59, 0x28, 0x35, 0xe8,
		0x00, 0xb2, 0x2c, 0x29, 0xc7, 0x3d, 0xbe, 0x94, 0x42, 0xaa, 0x36, 0xb0,
		0x17, 0x2c, 0x53, 0x28, 0x6f, 0xf2, 0xd9, 0xa9, 0xd0, 0xa7, 0xcc, 0x52,
		0x17, 0xc7, 0x1a, 0x07, 0x0e, 0xf4, 0x6e, 0x84, 0xfb, 0x13, 0xd1, 0xf9,
		0xf3, 0x54, 0x5c, 0x62, 0xf2, 0xe6, 0x27, 0xff, 0x37, 0xc4, 0xc0, 0x42,
		0xa7, 0x7e, 0x7e, 0x95, 0xd7, 0x5b, 0xfa, 0xd9, 0x71, 0x31, 0x4e, 0xfc,
		0xe1, 0x5f, 0x7b, 0xaf, 0xbe, 0x03, 0x00, 0x00, 0xff, 0xff, 0xb6, 0x49,
		0x6e, 0xbb, 0x96, 0x01, 0x00, 0x00,
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
