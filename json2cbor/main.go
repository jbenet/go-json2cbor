package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	j2c "github.com/jbenet/go-json2cbor"
)

var codec string
var usageStr = `json2cbor j2c - recode json to cbor
json2cbor c2j - recode cbor to json

note: reads newline delimited json. unescaped newlines in
json input may be a problem.

OPTIONS:
`

func init() {
	codecStr := strings.Join(j2c.Codecs, ", ")
	flag.StringVar(&codec, "codec", "ugorji", "cbor codec to use: "+codecStr)
	flag.Usage = usage
}

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stdout, err)
		os.Exit(1)
	}
}

func usage() {
	fmt.Fprintln(os.Stdout, usageStr)
	flag.PrintDefaults()
	os.Exit(0)
}

func run() error {
	flag.Parse()
	args := flag.Args()

	switch len(args) {
	default:
		usage()
		return nil
	case 1:
	}

	switch args[0] {
	default:
		usage()
		return nil
	case "j2c":
		return j2c.JsonToCbor(codec, os.Stdout, os.Stdin)
	case "c2j":
		return j2c.CborToJson(codec, os.Stdout, os.Stdin)
	}
}
