package utils

import (
	"runtime"
)

func UseMaxCpu() {
	// multiple cpus using
	runtime.GOMAXPROCS(runtime.NumCPU())
}
