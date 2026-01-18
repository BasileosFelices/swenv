package commands

import (
	"context"
	"fmt"
	"os"
	"sort"
	"strings"

	"github.com/BasileosFelices/swenv/internal/commands/common"
	"github.com/urfave/cli/v3"
)

type locatedEnvFiles struct {
	BaseEnvFileLocated bool
	FoundEnvFiles      []string
	FoundExampleFiles  []string
}

func PrettyPrintLocatedEnvFiles(located *locatedEnvFiles) string {
	var b strings.Builder

	inUse := located.BaseEnvFileLocated
	b.WriteString("Active .env: ")
	if inUse {
		b.WriteString("yes (.env present)\n")
	} else {
		b.WriteString("no\n")
	}

	b.WriteString("\nEnvironments:\n")
	if len(located.FoundEnvFiles) == 0 {
		b.WriteString("  (none)\n")
	} else {
		sort.Strings(located.FoundEnvFiles)
		for _, f := range located.FoundEnvFiles {
			env := strings.TrimPrefix(f, common.CustomEnvPrefix) // ".env.<env>"
			if env == "" {
				env = "(base)"
			}
			fmt.Fprintf(&b, "  - %s (%s)\n", env, f)
		}
	}

	b.WriteString("\nExamples:\n")
	if len(located.FoundExampleFiles) == 0 {
		b.WriteString("  (none)\n")
	} else {
		sort.Strings(located.FoundExampleFiles)
		for _, f := range located.FoundExampleFiles {
			env := strings.TrimSuffix(f, common.ExampleSuffix)    // drop ".example"
			env = strings.TrimPrefix(env, common.CustomEnvPrefix) // drop ".env."
			if env == "" || env == ".env" {
				env = "(base)"
			}
			fmt.Fprintf(&b, "  - %s (%s)\n", env, f)
		}
	}

	return b.String()
}

func ListEnvFiles(context.Context, *cli.Command) error {

	files, err := os.ReadDir(".")
	if err != nil {
		return cli.Exit("Failed to read current directory", 1)
	}

	located := &locatedEnvFiles{}

	for _, file := range files {
		name := file.Name()

		if file.IsDir() {
			continue
		}

		switch {
		case name == common.BaseEnvFile:
			located.BaseEnvFileLocated = true
		case strings.HasPrefix(name, common.CustomEnvPrefix) && strings.HasSuffix(name, common.ExampleSuffix):
			located.FoundExampleFiles = append(located.FoundExampleFiles, name)
		case strings.HasPrefix(name, common.CustomEnvPrefix):
			located.FoundEnvFiles = append(located.FoundEnvFiles, name)
		default:
			continue
		}
	}

	fmt.Print(PrettyPrintLocatedEnvFiles(located))

	return nil
}
