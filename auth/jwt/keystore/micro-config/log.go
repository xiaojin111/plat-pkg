package microconfig

import (
	mlog "github.com/jinmukeji/go-pkg/v2/log"
)

var (
	// log is the package global logger
	log *mlog.Logger
)

func init() {
	log = mlog.StandardLogger()
}
