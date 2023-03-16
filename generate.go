package main

//go:generate thrift -out pkg -recurse -gen go:package_prefix=github.com/bgould/tinygo-thrift-test/pkg/ idl/config.thrift
//go:generate gofmt -w pkg/config
