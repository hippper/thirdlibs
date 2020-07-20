package grmon

import (
	"sync"

	"github.com/luckyweiwei/base/utils"
)

type TGRMon struct {
	grmap *sync.Map
}

var grmon = &TGRMon{
	grmap: &sync.Map{},
}

func GetGRMon() *TGRMon {
	return grmon
}

func (s *TGRMon) addGR(name string) {
	s.grmap.Store(name, 1)
}

func (s *TGRMon) removeGR(name string) {
	s.grmap.Delete(name)
}

func (s *TGRMon) Go(name string, fn interface{}, args ...interface{}) {

	go func() {

		defer utils.CatchExceptionWithName(name)

		s.addGR(name)
		defer s.removeGR(name)

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

		s.addGR(name)
		defer s.removeGR(name)

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
