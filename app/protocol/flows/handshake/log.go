package handshake

import (
	"github.com/slyvex-core/slyvexd/infrastructure/logger"
	"github.com/slyvex-core/slyvexd/util/panics"
)

var log = logger.RegisterSubSystem("PROT")
var spawn = panics.GoroutineWrapperFunc(log)
