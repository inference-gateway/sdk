package sdk

import (
	"os"
	"os/exec"
	"path/filepath"
	"testing"

	require "github.com/stretchr/testify/require"
)

// TestExamplesBuild builds every standalone example module under examples/,
// mirroring `task build-examples`. The root build and test steps only compile
// the root module and skip nested modules, so a stale example go.mod or
// go.sum would go unnoticed; this test runs as part of `go test ./...`,
// which CI already runs, and fails when any example module no longer builds.
func TestExamplesBuild(t *testing.T) {
	matches, err := filepath.Glob(filepath.Join("examples", "*", "go.mod"))
	require.NoError(t, err)
	require.NotEmpty(t, matches, "no standalone example modules found under examples/")

	for _, gomod := range matches {
		dir := filepath.Dir(gomod)
		name := filepath.Base(dir)
		t.Run(name, func(t *testing.T) {
			cmd := exec.Command("go", "build", "-o", os.DevNull, "./...")
			cmd.Dir = dir
			out, err := cmd.CombinedOutput()
			require.NoError(t, err, "example module %s failed to build:\n%s", name, out)
		})
	}
}
