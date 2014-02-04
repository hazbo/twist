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
		0x41, 0x4f, 0x02, 0x31, 0x10, 0x85, 0xcf, 0xf6, 0x57, 0xbc, 0x70, 0x02,
		0x54, 0xd0, 0x2b, 0xc4, 0x93, 0xc1, 0x48, 0x82, 0x5e, 0xe4, 0xe8, 0xa5,
		0xb6, 0x83, 0x4c, 0x52, 0xba, 0xa4, 0x3b, 0x6c, 0x25, 0x86, 0xff, 0xee,
		0x74, 0x77, 0x55, 0xbc, 0x39, 0x49, 0x0f, 0xfd, 0x66, 0xde, 0xcc, 0x7b,
		0x66, 0x3a, 0x1e, 0x1b, 0x8c, 0xb1, 0xce, 0x5c, 0x0b, 0x9a, 0x9b, 0xc9,
		0xed, 0xb5, 0xa7, 0x46, 0x49, 0x81, 0x43, 0x37, 0xc2, 0xa3, 0x4d, 0xe9,
		0x88, 0x95, 0xcd, 0x89, 0xa2, 0xa3, 0xbe, 0xb1, 0x62, 0x47, 0xb1, 0xa6,
		0x19, 0x9e, 0x96, 0xeb, 0xf2, 0x2f, 0xef, 0xa1, 0x4a, 0x90, 0x2d, 0x61,
		0x73, 0x08, 0x01, 0xae, 0xda, 0x1f, 0x13, 0xbf, 0x6f, 0x05, 0x36, 0x7a,
		0x84, 0x6e, 0x1c, 0x1c, 0x37, 0x55, 0xda, 0x59, 0xe1, 0x2a, 0x5e, 0x61,
		0x1f, 0xc8, 0x2a, 0x6b, 0x98, 0x72, 0xab, 0x5b, 0x2d, 0xef, 0x17, 0xcf,
		0x2f, 0x8b, 0xb2, 0x6a, 0xc3, 0x81, 0x94, 0x59, 0x41, 0xb6, 0x35, 0xbc,
		0x3a, 0x4b, 0xfc, 0x76, 0x10, 0xf2, 0xc8, 0x2c, 0x5b, 0xed, 0x70, 0x8d,
		0xba, 0x3a, 0x24, 0x47, 0x7a, 0xc8, 0xd3, 0x44, 0x35, 0x53, 0x63, 0x1a,
		0x9b, 0x40, 0x9e, 0x45, 0x7d, 0xdc, 0xe1, 0xd3, 0x18, 0x68, 0xed, 0x13,
		0x47, 0xc1, 0x4c, 0x4d, 0x45, 0x57, 0xce, 0x0e, 0x85, 0x3e, 0x64, 0xa4,
		0x5d, 0xf4, 0xd5, 0x0e, 0x74, 0x74, 0xde, 0xc2, 0xd3, 0x99, 0x30, 0xc4,
		0xff, 0x49, 0x71, 0x89, 0xc1, 0x6b, 0x1c, 0xfc, 0xdd, 0x90, 0x13, 0x0b,
		0x9d, 0xeb, 0x4b, 0xaa, 0x68, 0x77, 0xf4, 0xbd, 0xe3, 0xa2, 0x9d, 0xf8,
		0xc5, 0x3f, 0xea, 0x53, 0x9f, 0x45, 0x63, 0x74, 0x79, 0xe6, 0x5f, 0x01,
		0x00, 0x00, 0xff, 0xff, 0x51, 0x8d, 0x4f, 0x37, 0xa6, 0x01, 0x00, 0x00,
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