package service

import "github.com/qisst/ms-nadra-verification/conf"

type Container struct {
	Config        *conf.GbeConfig
	SampleService SampleService
}

// create container by intiating all required services
func NewServiceContainer() *Container {
	config := conf.GetConfig()
	sampleService := NewSampleService()
	return &Container{
		Config:        config,
		SampleService: sampleService,
	}
}
