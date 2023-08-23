package conf

import (
	"testing"
)

func TestGetConfigPanic(t *testing.T) {

	t.Run("Should not load config", func(t *testing.T) {

		defer func() {
			r := recover()
			if r == nil {
				t.Errorf("The code did not panic")
			}
		}()

		SetConfFilePath("in valid path")
		SetConfFileName("invalid name")
		GetConfig()

	})

}

func TestGetConfig(t *testing.T) {

	t.Run("Should load config", func(t *testing.T) {

		SetConfFilePath("../")
		GetConfig()

	})
}
