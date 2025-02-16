package common

import (
	"fmt"
	"github.com/slyvex-core/slyvexd/domain/dagconfig"
	"os"
	"sync/atomic"
	"syscall"
	"testing"
)

// RunSlyvexdForTesting runs slyvexd for testing purposes
func RunSlyvexdForTesting(t *testing.T, testName string, rpcAddress string) func() {
	appDir, err := TempDir(testName)
	if err != nil {
		t.Fatalf("TempDir: %s", err)
	}

	slyvexdRunCommand, err := StartCmd("SLYVEXD",
		"slyvexd",
		NetworkCliArgumentFromNetParams(&dagconfig.DevnetParams),
		"--appdir", appDir,
		"--rpclisten", rpcAddress,
		"--loglevel", "debug",
	)
	if err != nil {
		t.Fatalf("StartCmd: %s", err)
	}
	t.Logf("slyvexd started with --appdir=%s", appDir)

	isShutdown := uint64(0)
	go func() {
		err := slyvexdRunCommand.Wait()
		if err != nil {
			if atomic.LoadUint64(&isShutdown) == 0 {
				panic(fmt.Sprintf("slyvexd closed unexpectedly: %s. See logs at: %s", err, appDir))
			}
		}
	}()

	return func() {
		err := slyvexdRunCommand.Process.Signal(syscall.SIGTERM)
		if err != nil {
			t.Fatalf("Signal: %s", err)
		}
		err = os.RemoveAll(appDir)
		if err != nil {
			t.Fatalf("RemoveAll: %s", err)
		}
		atomic.StoreUint64(&isShutdown, 1)
		t.Logf("slyvexd stopped")
	}
}
