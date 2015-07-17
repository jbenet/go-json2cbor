package main

import (
	"errors"
	"fmt"
	"os"

	j2c "github.com/jbenet/go-json2cbor"
)

var usage = `json2cbor j2c - recode json to cbor
json2cbor c2j - recode cbor to json

special:
  json2cbor j2cs - recode json to cbor (using encoding/json)

note: reads newline delimited json. unescaped newlines in
json input may be a problem.
`

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stdout, err)
		os.Exit(1)
	}
}

func run() error {
	switch len(os.Args) {
	default:
		return errors.New(usage)
	case 2:
	}

	switch os.Args[1] {
	default:
		return errors.New(usage)
	case "j2cs":
		return j2c.JsonToCborS(os.Stdout, os.Stdin)
	case "j2c":
		return j2c.JsonToCbor(os.Stdout, os.Stdin)
	case "c2j":
		return j2c.CborToJson(os.Stdout, os.Stdin)
	}
}
