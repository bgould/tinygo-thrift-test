package main

import (
	"context"

	"github.com/bgould/tinygo-thrift-test/pkg/config"
	"github.com/bgould/tinygo-thrift-test/pkg/thriftutil"
)

var exampleConfig = &config.Root{
	Wifi: &config.Wifi{
		Ssid:           "TestSSID",
		Passphrase:     "MyPassphrase",
		TimeoutSeconds: 5,
	},
}

func main() {

	ctx := context.Background()

	thriftutil.WriteFile(ctx, "example_config.thrift", exampleConfig)

}
