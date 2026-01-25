package commands_test

import (
	"context"
	"io"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/BasileosFelices/swenv/internal/commands"
	"github.com/BasileosFelices/swenv/internal/commands/common"
	"github.com/urfave/cli/v3"
)

func withTempDir(t *testing.T) string {
	t.Helper()

	dir := t.TempDir()
	cwd, err := os.Getwd()
	if err != nil {
		t.Fatalf("getwd: %v", err)
	}

	if err := os.Chdir(dir); err != nil {
		t.Fatalf("chdir: %v", err)
	}

	t.Cleanup(func() {
		_ = os.Chdir(cwd)
	})

	return dir
}

func writeFile(t *testing.T, name string, content string) {
	t.Helper()

	if err := os.WriteFile(name, []byte(content), 0o644); err != nil {
		t.Fatalf("write %s: %v", name, err)
	}
}

func readFile(t *testing.T, name string) string {
	t.Helper()

	data, err := os.ReadFile(name)
	if err != nil {
		t.Fatalf("read %s: %v", name, err)
	}
	return string(data)
}

func runCLI(t *testing.T, args ...string) (string, error) {
	t.Helper()

	cmd := commands.NewRootCommand()
	var runErr error
	cmd.ExitErrHandler = func(ctx context.Context, cmd *cli.Command, err error) {
		runErr = err
	}
	cmd.ErrWriter = io.Discard
	stdout := captureStdout(t)

	_ = cmd.Run(context.Background(), append([]string{"swenv"}, args...))
	output := stdout()

	return output, runErr
}

func captureStdout(t *testing.T) func() string {
	t.Helper()

	old := os.Stdout
	r, w, err := os.Pipe()
	if err != nil {
		t.Fatalf("pipe: %v", err)
	}

	os.Stdout = w

	return func() string {
		_ = w.Close()
		os.Stdout = old

		out, _ := io.ReadAll(r)
		return string(out)
	}
}

func TestSwitchEnvFile_OnlyBaseAndBackupChanged(t *testing.T) {
	withTempDir(t)

	writeFile(t, common.BaseEnvFile, "BASE=1\n")
	writeFile(t, ".env.dev", "DEV=1\n")
	writeFile(t, ".env.prod", "PROD=1\n")
	writeFile(t, ".env.stuff", "STUFF=1\n")
	writeFile(t, ".env.dev.example", "EXAMPLE=1\n")

	before := map[string]string{
		common.BaseEnvFile: readFile(t, common.BaseEnvFile),
		".env.dev":         readFile(t, ".env.dev"),
		".env.prod":        readFile(t, ".env.prod"),
		".env.stuff":       readFile(t, ".env.stuff"),
		".env.dev.example": readFile(t, ".env.dev.example"),
	}

	if _, err := runCLI(t, "switch", "dev"); err != nil {
		t.Fatalf("switch failed: %v", err)
	}

	if got := readFile(t, common.BaseEnvFile); got != "DEV=1\n" {
		t.Fatalf(".env content mismatch: %q", got)
	}

	if got := readFile(t, common.BackupEnvFile); got != "BASE=1\n" {
		t.Fatalf("backup content mismatch: %q", got)
	}

	for name, content := range before {
		if name == common.BaseEnvFile || name == common.BackupEnvFile {
			continue
		}
		if got := readFile(t, name); got != content {
			t.Fatalf("%s changed unexpectedly: %q", name, got)
		}
	}
}

func TestSwitchEnvFile_MissingNameReturnsError(t *testing.T) {
	withTempDir(t)

	_, err := runCLI(t, "switch")
	if err == nil {
		t.Fatalf("expected error, got nil")
	}

	if !strings.Contains(err.Error(), "Please provide the name of the env environment to switch to") {
		t.Fatalf("unexpected error message: %v", err)
	}
}

func TestSwitchEnvFile_MissingSourceReturnsError(t *testing.T) {
	withTempDir(t)

	_, err := runCLI(t, "switch", "missing")
	if err == nil {
		t.Fatalf("expected error, got nil")
	}

	if !strings.Contains(err.Error(), "Env file [.env.missing] does not exist") {
		t.Fatalf("unexpected error message: %v", err)
	}

	if _, statErr := os.Stat(common.BackupEnvFile); !os.IsNotExist(statErr) {
		t.Fatalf("unexpected backup file created: %v", statErr)
	}
}

func TestSwitchEnvFile_DoesNotAlterUnrelatedFiles(t *testing.T) {
	withTempDir(t)

	writeFile(t, common.BaseEnvFile, "BASE=1\n")
	writeFile(t, ".env.dev", "DEV=1\n")
	writeFile(t, ".env.other", "OTHER=1\n")

	originalOther := readFile(t, ".env.other")

	if _, err := runCLI(t, "switch", "dev"); err != nil {
		t.Fatalf("switch failed: %v", err)
	}

	if got := readFile(t, ".env.other"); got != originalOther {
		t.Fatalf(".env.other changed unexpectedly: %q", got)
	}

	if _, err := os.Stat(filepath.Join(".", ".env.other")); err != nil {
		t.Fatalf(".env.other missing: %v", err)
	}
}
