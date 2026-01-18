package main

import (
	"context"
	"log"
	"os"

	"github.com/urfave/cli/v3"

	"github.com/BasileosFelices/swenv/internal/commands"
)

func main() {
	cmd := &cli.Command{
		Name:   "swenv",
		Usage:  "Quickly manage multiple .env files in a project",
		Action: commands.ListEnvFiles,
		Commands: []*cli.Command{
			{
				Name:    "switch",
				Aliases: []string{"sw"},
				Usage:   "Switch to a specified .env file",
				Arguments: []cli.Argument{
					&cli.StringArg{
						Name:      "name",
						UsageText: "Name of the desired env enviroment",
						Config:    cli.StringConfig{TrimSpace: true},
					},
				},
				Action: commands.SwitchEnvFile,
			},
			{
				Name:    "list",
				Aliases: []string{"ls"},
				Usage:   "List all available .env files",
				Action:  commands.ListEnvFiles,
			},
		},
	}

	if err := cmd.Run(context.TODO(), os.Args); err != nil {
		log.Fatal(err)
	}
}
