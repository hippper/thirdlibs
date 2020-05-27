package logger

import (
	"errors"
	"fmt"
	"testing"
)

func TestLogger(t *testing.T) {
	log := InitLog()

	fmt.Printf("%+v", log)

	log.Error(errors.New("test"))
}
