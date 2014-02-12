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
		0x31, 0x4f, 0xf3, 0x30, 0x10, 0x86, 0xe7, 0x2f, 0xbf, 0xe2, 0x55, 0xa7,
		0x36, 0x1f, 0xb4, 0xb0, 0xb6, 0x62, 0x42, 0x45, 0x54, 0x0a, 0x2c, 0x74,
		0x64, 0x31, 0xf6, 0x85, 0x9c, 0xe4, 0x3a, 0x91, 0x73, 0x89, 0xa9, 0x50,
		0xff, 0x3b, 0x76, 0x92, 0x52, 0xd8, 0x88, 0x94, 0xc1, 0xcf, 0xdd, 0x73,
		0x77, 0xef, 0x2a, 0xcf, 0x33, 0xe4, 0xd8, 0x07, 0x6e, 0x05, 0xfd, 0xcd,
		0xf2, 0xf6, 0xda, 0x50, 0x1f, 0x49, 0x82, 0x73, 0xbd, 0xc0, 0xa3, 0xf2,
		0xfe, 0x88, 0x42, 0x05, 0x4f, 0x4e, 0xd3, 0x54, 0x28, 0x58, 0x93, 0x6b,
		0x69, 0x8d, 0xa7, 0xdd, 0x3e, 0xbd, 0xd3, 0xff, 0x50, 0x7b, 0x48, 0x45,
		0x28, 0x3b, 0x6b, 0xa1, 0xeb, 0xe6, 0xe8, 0xf9, 0xbd, 0x12, 0x28, 0x67,
		0x60, 0xc7, 0x76, 0xb0, 0x2b, 0x6b, 0x7f, 0x50, 0xc2, 0xb5, 0xbb, 0x42,
		0x63, 0x49, 0x45, 0xd6, 0x33, 0x85, 0xc1, 0x2b, 0x76, 0xf7, 0xdb, 0xe7,
		0x97, 0x6d, 0x1a, 0x55, 0xb2, 0xa5, 0xc8, 0x94, 0x20, 0xa8, 0x16, 0x26,
		0x5e, 0xe6, 0xf9, 0xad, 0x13, 0x32, 0x08, 0x2c, 0x55, 0xac, 0x70, 0x8b,
		0xb6, 0xee, 0xbc, 0xa6, 0xb8, 0xc8, 0xd0, 0x32, 0x3a, 0xab, 0x2c, 0xeb,
		0x95, 0x07, 0x19, 0x96, 0x78, 0xc7, 0x1d, 0x3e, 0x33, 0xc4, 0xaf, 0xf1,
		0xec, 0x04, 0xeb, 0x78, 0x93, 0xd3, 0x69, 0xeb, 0x5c, 0xe8, 0x43, 0x16,
		0x53, 0xf1, 0xbb, 0x61, 0xa4, 0x9b, 0x01, 0x9e, 0xb2, 0x8b, 0x68, 0xdd,
		0xdf, 0x54, 0xfc, 0xc7, 0xec, 0xd5, 0xcd, 0x7e, 0x4f, 0x08, 0x9e, 0x85,
		0x7e, 0xfa, 0x29, 0x94, 0x53, 0x07, 0x3a, 0xcf, 0xf8, 0x37, 0x74, 0x5c,
		0xf0, 0xd9, 0x3e, 0x4d, 0x49, 0x62, 0x88, 0x31, 0xcd, 0xe6, 0x2b, 0x00,
		0x00, 0xff, 0xff, 0xfc, 0x6f, 0x74, 0xb6, 0xa3, 0x01, 0x00, 0x00,
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
