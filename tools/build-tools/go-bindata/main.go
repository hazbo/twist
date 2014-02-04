// This work is subject to the CC0 1.0 Universal (CC0 1.0) Public Domain Dedication
// license. Its contents can be found at:
// http://creativecommons.org/publicdomain/zero/1.0/

package main

import (
	"flag"
	"fmt"
	"../../../vendor/go-bindata"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	cfg := parseArgs()
	err := bindata.Translate(cfg)

	if err != nil {
		fmt.Fprintf(os.Stderr, "bindata: %v\n", err)
		os.Exit(1)
	}
}

// parseArgs create s a new, filled configuration instance
// by reading and parsing command line options.
//
// This function exits the program with an error, if
// any of the command line options are incorrect.
func parseArgs() *bindata.Config {
	var version bool

	c := bindata.NewConfig()

	flag.Usage = func() {
		fmt.Printf("Usage: %s [options] <input directories>\n\n", os.Args[0])
		flag.PrintDefaults()
	}

	flag.BoolVar(&c.Debug, "debug", c.Debug, "Do not embed the assets, but provide the embedding API. Contents will still be loaded from disk.")
	flag.StringVar(&c.Tags, "tags", c.Tags, "Optional set of uild tags to include.")
	flag.StringVar(&c.Prefix, "prefix", c.Prefix, "Optional path prefix to strip off asset names.")
	flag.StringVar(&c.Package, "pkg", c.Package, "Package name to use in the generated code.")
	flag.BoolVar(&c.NoMemCopy, "nomemcopy", c.NoMemCopy, "Use a .rodata hack to get rid of unnecessary memcopies. Refer to the documentation to see what implications this carries.")
	flag.BoolVar(&c.NoCompress, "nocompress", c.NoCompress, "Assets will *not* be GZIP compressed when this flag is specified.")
	flag.StringVar(&c.Output, "o", c.Output, "Optional name of the output file to be generated.")
	flag.BoolVar(&version, "version", false, "Displays version information.")
	flag.Parse()

	if version {
		fmt.Printf("%s\n", Version())
		os.Exit(0)
	}

	// Make sure we have input paths.
	if flag.NArg() == 0 {
		fmt.Fprintf(os.Stderr, "Missing <input dir>\n\n")
		flag.Usage()
		os.Exit(1)
	}

	// Create input configurations.
	c.Input = make([]bindata.InputConfig, flag.NArg())
	for i := range c.Input {
		c.Input[i] = parseInput(flag.Arg(i))
	}

	return c
}

// parseRecursive determines whether the given path has a recrusive indicator and
// returns a new path with the recursive indicator chopped off if it does.
//
//  ex:
//      /path/to/foo/...    -> (/path/to/foo, true)
//      /path/to/bar        -> (/path/to/bar, false)
func parseInput(path string) bindata.InputConfig {
	if strings.HasSuffix(path, "/...") {
		return bindata.InputConfig{
			Path:      filepath.Clean(path[:len(path)-4]),
			Recursive: true,
		}
	} else {
		return bindata.InputConfig{
			Path:      filepath.Clean(path),
			Recursive: false,
		}
	}

}
