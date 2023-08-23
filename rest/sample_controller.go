package rest

type sampleController struct {
}

type SampleController interface {
}

func NewSampleController() SampleController {
	return sampleController{}
}
