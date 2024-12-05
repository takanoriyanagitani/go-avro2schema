package avsc2writer

import (
	"bufio"
	"context"
	"io"
	"os"
	"strings"

	aa "github.com/takanoriyanagitani/go-avro2schema"
	util "github.com/takanoriyanagitani/go-avro2schema/util"
)

func AvscToWriter(w io.Writer) func(aa.AvroSchema) error {
	return func(s aa.AvroSchema) error {
		var bw *bufio.Writer = bufio.NewWriter(w)
		defer bw.Flush()

		var r io.Reader = strings.NewReader(string(s))
		_, e := io.Copy(bw, r)
		return e
	}
}

func WriterToSchemaSink(w io.Writer) func(aa.AvroSchema) util.IO[util.Void] {
	return func(s aa.AvroSchema) util.IO[util.Void] {
		return func(_ context.Context) (util.Void, error) {
			return util.Empty, AvscToWriter(w)(s)
		}
	}
}

var SchemaToStdoutSink func(
	aa.AvroSchema,
) util.IO[util.Void] = WriterToSchemaSink(os.Stdout)
