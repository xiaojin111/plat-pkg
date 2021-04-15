package microconfig

import (
	mlog "gitee.com/jt-heath/go-pkg/v2/log"
)

var (
	// log is the package global logger
	log *mlog.Logger
)

func init() {
	log = mlog.StandardLogger()
}
