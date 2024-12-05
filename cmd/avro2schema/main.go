package main

import (
	"context"
	"log"

	aa "github.com/takanoriyanagitani/go-avro2schema"
	util "github.com/takanoriyanagitani/go-avro2schema/util"

	rh "github.com/takanoriyanagitani/go-avro2schema/avro/dec/reader/hamba"
	sw "github.com/takanoriyanagitani/go-avro2schema/avro/schema/sink/writer"
)

var stdin2schema util.IO[aa.AvroSchema] = rh.StdinToSchema
var schema2stdout func(aa.AvroSchema) util.IO[util.Void] = sw.SchemaToStdoutSink

var stdin2schema2stdout util.IO[util.Void] = util.Bind(
	stdin2schema,
	schema2stdout,
)

func sub(ctx context.Context) error {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	_, e := stdin2schema2stdout(ctx)
	return e
}

func main() {
	e := sub(context.Background())
	if nil != e {
		log.Printf("%v\n", e)
	}
}
