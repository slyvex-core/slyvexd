package main

import (
	"github.com/slyvex-core/slyvexd/infrastructure/logger"
	"github.com/slyvex-core/slyvexd/util/panics"
)

var (
	backendLog = logger.NewBackend()
	log        = backendLog.Logger("RPIC")
	spawn      = panics.GoroutineWrapperFunc(log)
)
