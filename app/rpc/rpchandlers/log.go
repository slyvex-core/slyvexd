package rpchandlers

import (
	"github.com/slyvex-core/slyvexd/infrastructure/logger"
	"github.com/slyvex-core/slyvexd/util/panics"
)

var log = logger.RegisterSubSystem("RPCS")
var spawn = panics.GoroutineWrapperFunc(log)
