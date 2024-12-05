package avro2avsc

import (
	"io"
	"os"

	ho "github.com/hamba/avro/v2/ocf"

	aa "github.com/takanoriyanagitani/go-avro2schema"
	util "github.com/takanoriyanagitani/go-avro2schema/util"
)

func ReaderToDecoder(r io.Reader) (*ho.Decoder, error) {
	return ho.NewDecoder(r)
}

func DecoderToSchema(d *ho.Decoder) (aa.AvroSchema, error) {
	return aa.AvroSchema(d.Schema().String()), nil
}

var StdinToDecoder util.IO[*ho.Decoder] = util.Bind(
	util.Of(io.Reader(os.Stdin)),
	util.Lift(ReaderToDecoder),
)

var StdinToSchema util.IO[aa.AvroSchema] = util.Bind(
	StdinToDecoder,
	util.Lift(DecoderToSchema),
)
