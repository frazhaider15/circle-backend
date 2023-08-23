package service

type sampleService struct {
}

type SampleService interface {
}

func NewSampleService() SampleService {
	return sampleService{}
}
