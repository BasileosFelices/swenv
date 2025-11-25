package commands

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/urfave/cli/v3"
)

const (
	baseEnvFile     = ".env"
	customEnvPrefix = ".env."
)

func ListEnvFiles(context.Context, *cli.Command) error {

	files, err := os.ReadDir(".")
	if err != nil {
		return cli.Exit("Failed to read current directory", 1)
	}

	anyFound := false
	for _, file := range files {
		name := file.Name()
		if !file.IsDir() && (strings.HasPrefix(name, customEnvPrefix) || name == baseEnvFile) {
			envName := strings.TrimPrefix(file.Name(), customEnvPrefix)
			fmt.Printf("  - %s (%s)\n", envName, name)
			anyFound = true
		}
	}
	if !anyFound {
		fmt.Println("No .env files found in the current directory.")
	}
	return nil
}
