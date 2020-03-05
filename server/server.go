package server

import (
	"github.com/error2215/simple_mongodb/server/api"
	"github.com/error2215/simple_mongodb/server/api/grpc"
	"github.com/error2215/simple_mongodb/server/api/rest"
	"github.com/error2215/simple_mongodb/server/config"

	log "github.com/sirupsen/logrus"

	"sync"
)

func Start() {
	restPort := config.GlobalConfig.RESTPort
	grpcPort := config.GlobalConfig.GRPCPort

	log.WithFields(log.Fields{
		"grpcPort": grpcPort,
		"restPort": restPort,
	}).Info("Launching API server")

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		g := &grpc.Server{}
		start(g)
	}()
	go func() {
		defer wg.Done()
		g := &rest.Server{}
		start(g)
	}()

	wg.Wait()
}

func start(api api.API) {
	api.Start()
}
