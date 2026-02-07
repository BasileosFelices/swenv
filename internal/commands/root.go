package commands

import (
	"context"

	"github.com/urfave/cli/v3"
)

// NewRootCommand builds the CLI root command for swenv.
func NewRootCommand() *cli.Command {
	return &cli.Command{
		Name:  "swenv",
		Usage: "Quickly manage multiple .env files in a project",
		Action: func(ctx context.Context, cmd *cli.Command) error {
			// If arguments are provided, default is switch command
			if cmd.Args().Len() > 0 {
				return SwitchEnvFile(ctx, cmd)
			}
			return ListEnvFiles(ctx, cmd)
		},
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
				Action: SwitchEnvFile,
			},
			{
				Name:    "list",
				Aliases: []string{"ls"},
				Usage:   "List all available .env files",
				Action:  ListEnvFiles,
			},
		},
	}
}
