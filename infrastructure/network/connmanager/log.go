package connmanager

import (
	"github.com/slyvex-core/slyvexd/infrastructure/logger"
	"github.com/slyvex-core/slyvexd/util/panics"
)

var log = logger.RegisterSubSystem("CMGR")
var spawn = panics.GoroutineWrapperFunc(log)
