package commands

import "github.com/urfave/cli/v3"

// NewRootCommand builds the CLI root command for swenv.
func NewRootCommand() *cli.Command {
	return &cli.Command{
		Name:   "swenv",
		Usage:  "Quickly manage multiple .env files in a project",
		Action: ListEnvFiles,
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