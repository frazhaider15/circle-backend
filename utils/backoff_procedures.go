package utils

import (
	"errors"
	"time"
)

type BackOff struct {
	Retried    int
	MaxRetries int
	Tick       time.Time
	RetryTime  time.Duration
	Operation  func(...interface{}) error
}

var (
	ErrMaxRetryLimit = errors.New("MAX_RETRY_LIMIT")
	ErrQuit          = errors.New("QUITBACKOFF")
)

func NewBackOff(op func(params ...interface{}) error, retryTime time.Duration) BackOff {
	return BackOff{
		Operation: op,
		RetryTime: retryTime,
	}
}

//ContinousBackOff try to do operation if it fail to maximum reties provided.
// -1 for max number of tries. (2147483646)
func (b BackOff) ContinousBackOff(maxTries int, params ...interface{}) error {
	if maxTries == -1 {
		b.MaxRetries = 2147483646
	}

	for {
		if b.Retried >= b.MaxRetries {
			break
		}

		err := b.Operation(params)
		if errors.Is(err, ErrQuit) {
			return ErrQuit
		}
		if err == nil {
			return nil
		}
		b.Retried = b.Retried + 1

		<-time.NewTicker(b.RetryTime * time.Duration(b.Retried)).C
	}

	return ErrMaxRetryLimit
}
