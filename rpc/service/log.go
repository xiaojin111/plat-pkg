package service

import mlog "github.com/jinmukeji/go-pkg/v2/log"

var (
	// log is the package global logger
	log = mlog.StandardLogger()
)

func Logger() *mlog.Logger {
	return mlog.StandardLogger()
}
