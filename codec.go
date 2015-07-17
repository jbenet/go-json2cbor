package json2cbor

import (
	"errors"
	"io"

	dwcbor "github.com/DamnWidget/cbor"
	ugorji "github.com/ugorji/go/codec"
)

var ErrCodecNotSupported = errors.New("codec not supported")
var Codecs = []string{"dw", "ugorji"}

func JsonToCbor(codec string, w io.Writer, r io.Reader) error {
	switch codec {
	case "dw":
		return DWJsonToCbor(w, r)
	case "ugorji":
		return UgorjiJsonToCbor(w, r)
	default:
		return ErrCodecNotSupported
	}
}

func CborToJson(codec string, w io.Writer, r io.Reader) error {
	switch codec {
	case "dw":
		return DWCborToJson(w, r)
	case "ugorji":
		return UgorjiCborToJson(w, r)
	default:
		return ErrCodecNotSupported
	}
}

func DWJsonToCbor(w io.Writer, r io.Reader) error {
	dec := ugorji.NewDecoder(r, &ugorji.JsonHandle{})
	enc := dwcbor.NewEncoder(w)
	return EDJsonToCbor(w, enc, dec)
}

func DWCborToJson(w io.Writer, r io.Reader) error {
	dec := dwcbor.NewDecoder(r)
	enc := ugorji.NewEncoder(w, &ugorji.JsonHandle{})
	return EDCborToJson(enc, dec)
}

func UgorjiJsonToCbor(w io.Writer, r io.Reader) error {
	dec := ugorji.NewDecoder(r, &ugorji.JsonHandle{})
	enc := ugorji.NewEncoder(w, &ugorji.CborHandle{})
	return EDJsonToCbor(w, enc, dec)
}

func UgorjiCborToJson(w io.Writer, r io.Reader) error {
	dec := ugorji.NewDecoder(r, &ugorji.CborHandle{})
	enc := ugorji.NewEncoder(w, &ugorji.JsonHandle{})
	return EDCborToJson(enc, dec)
}

func EDCborToJson(enc Encoder, dec Decoder) error {
	var o interface{}
	for {
		if err := dec.Decode(&o); err == io.EOF {
			return nil
		} else if err != nil {
			return err
		}

		if err := enc.Encode(&o); err != nil {
			return err
		}
	}
}

func EDJsonToCbor(w io.Writer, enc Encoder, dec Decoder) error {
	var o interface{}
	for {
		if err := dec.Decode(&o); err == io.EOF {
			return nil
		} else if err != nil {
			return err
		}

		if err := enc.Encode(&o); err != nil {
			return err
		}
		if _, err := w.Write([]byte("\n")); err != nil {
			return err
		}
	}
}

type Encoder interface {
	Encode(interface{}) error
}

type Decoder interface {
	Decode(interface{}) error
}
