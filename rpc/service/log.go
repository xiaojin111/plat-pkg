package service

import mlog "gitee.com/jt-heath/go-pkg/v2/log"

var (
	// log is the package global logger
	log = mlog.StandardLogger()
)

func Logger() *mlog.Logger {
	return mlog.StandardLogger()
}
