package json2cbor

import (
	"encoding/json"
	"io"

	cbor "github.com/gonuts/cbor"
	codec "github.com/ugorji/go/codec"
)

func JsonToCborS(w io.Writer, r io.Reader) error {
	var o interface{}
	cenc := cbor.NewEncoder(w)
	jdec := json.NewDecoder(r)

	for {
		if err := jdec.Decode(&o); err == io.EOF {
			return nil
		} else if err != nil {
			return err
		}

		if err := cenc.Encode(&o); err != nil {
			return err
		}
	}
}

func JsonToCbor(w io.Writer, r io.Reader) error {
	var o interface{}
	cenc := codec.NewEncoder(w, &codec.CborHandle{})
	jdec := codec.NewDecoder(r, &codec.JsonHandle{})

	for {
		if err := jdec.Decode(&o); err == io.EOF {
			return nil
		} else if err != nil {
			return err
		}

		if err := cenc.Encode(&o); err != nil {
			return err
		}
	}
}

func CborToJson(w io.Writer, r io.Reader) error {
	var o interface{}
	jenc := codec.NewEncoder(w, &codec.JsonHandle{})
	cdec := codec.NewDecoder(r, &codec.CborHandle{})

	for {
		if err := cdec.Decode(&o); err == io.EOF {
			return nil
		} else if err != nil {
			return err
		}

		if err := jenc.Encode(&o); err != nil {
			return err
		}
		if _, err := w.Write([]byte("\n")); err != nil {
			return err
		}
	}
}
