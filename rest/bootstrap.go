package rest

import (
	"github.com/qisst/ms-nadra-verification/logger"
	"github.com/qisst/ms-nadra-verification/service"
)

// StartServer initate server
func StartServer(container *service.Container) *HttpServer {

	//Inject services instance from ServiceContainer
	sampleController := NewSampleController()

	httpServer := NewHttpServer(
		container.Config.RestServer.Addr,
	)

	//Inject controller instance to server
	httpServer.sampleController = sampleController

	go httpServer.Start()
	logger.Instance().Info("rest server ok")
	return httpServer
}
