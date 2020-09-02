package grmon

import (
	"github.com/luckyweiwei/base/utils"
)

type TGRMon struct {
}

var grmon = &TGRMon{}

func GetGRMon() *TGRMon {
	return grmon
}

func (s *TGRMon) Go(name string, fn interface{}, args ...interface{}) {

	go func() {
		defer utils.CatchExceptionWithName(name)

		if len(args) == 0 {
			f := fn.(func())
			f()
		} else {
			f := fn.(func(args ...interface{}))
			f(args...)
		}
	}()
}

func (s *TGRMon) GoLoop(name string, fn interface{}, args ...interface{}) {

	go func() {
		defer utils.CatchExceptionWithName(name)

		for {
			if len(args) == 0 {
				f := fn.(func())
				f()
			} else {
				f := fn.(func(args ...interface{}))
				f(args...)
			}
		}
	}()
}
