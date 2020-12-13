package client

import (
	"github.com/micro/go-micro/v2/config"
	"github.com/micro/go-micro/v2/util/log"
	grpcConfig "github.com/micro/go-plugins/config/source/grpc/v2"
)

func LoadSource(address string, name string) (config.Config, error) {
	// create new client
	source := grpcConfig.NewSource(
		grpcConfig.WithAddress(address),
		grpcConfig.WithPath("/"+name),
	)

	// create new config
	conf, _ := config.NewConfig()

	// load the client into config
	if err := conf.Load(source); err != nil {
		return nil, err
	}
	return conf, nil
}

func WatchSource(conf config.Config) {
	// watch the config for changes
	watcher, err := conf.Watch()
	if err != nil {
		log.Fatal(err)
	}

	log.Logf("Watching for changes ...")

	for {
		v, err := watcher.Next()
		if err != nil {
			log.Fatal(err)
		}

		log.Logf("Watching for changes: %v", string(v.Bytes()))
	}
}
