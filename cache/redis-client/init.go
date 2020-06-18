package redisclient

import "github.com/luckyweiwei/base/logger"

var Log *logger.Logger = nil

func init() {
	Log = logger.Log
}
