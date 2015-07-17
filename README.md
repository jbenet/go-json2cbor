# json2cbor - recoder in go

This is a simple json2cbor recoder. All the heavy lifting is done
by `encoding/json` and [github.com/ugorji/go/codec](https://github.com/ugorji/go/tree/master/codec#readme).

## Usage

```go
package main

import j2c "github.com/jbenet/go-json2cbor"


func main() {
  switch os.Args[1] {
  default:
    return errors.New("arg must be j2c or c2j")
  case "j2c":
    return j2c.JsonToCbor(os.Stdout, os.Stdin)
  case "c2j":
    return j2c.CborToJson(os.Stdout, os.Stdin)
  }
}
```

## `json2cbor` tool

```
go get github.com/jbenet/go-json2cbor/json2cbor
```

convert json to cbor

```
cat a_bunch_of.json | json2cbor j2c >a_bunch_of.cbor
cat a_bunch_of.cbor | json2cbor c2j >a_bunch_of.json
```

outputs [newline-delimited json](//github.com/maxogden/ndjson).

Test codec with:

```
random-json -o | json2cbor j2c | json2cbor c2j
```

Test JSON->CBOR->JSON roundtripping with:

```
random-json -o -c 10 | tee a.json | json2cbor j2c | json2cbor c2j >b.json && diff a.json b.json
```
