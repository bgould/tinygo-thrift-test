package thriftutil

import (
	"context"
	"os"

	"github.com/apache/thrift/lib/go/thrift"
)

// var (
// 	FS os.Filesystem
// )

func ReadFile(ctx context.Context, path string, obj thrift.TStruct) error {
	f, err := os.OpenFile(path, os.O_RDONLY, 0664)
	if err != nil {
		return err
	}
	defer f.Close()
	tr := thrift.NewStreamTransportR(f)
	pr := thrift.NewTBinaryProtocol(tr, false, true)
	if err := obj.Read(ctx, pr); err != nil {
		return err
	}
	return nil
}

func WriteFile(ctx context.Context, path string, obj thrift.TStruct) error {
	f, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE, 0664)
	if err != nil {
		return err
	}
	defer f.Close()
	tr := thrift.NewStreamTransportW(f)
	defer tr.Flush(nil) // FIXME: shouldn't really pass nil for context
	pr := thrift.NewTBinaryProtocol(tr, false, true)
	if err := obj.Write(ctx, pr); err != nil {
		return err
	}
	return nil
}
