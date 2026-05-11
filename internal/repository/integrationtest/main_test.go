package integrationtest

import (
	"log/slog"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	if os.Getenv("INTEGRATION_TEST") == "" {
		return
	}
	slog.Info("Running integration tests in repository package...")

	exitCode := m.Run()
	os.Exit(exitCode)
}
