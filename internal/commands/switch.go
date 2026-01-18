package commands

import (
	"context"
	"fmt"

	"github.com/urfave/cli/v3"
)

func SwitchEnvFile(ctx context.Context, cmd *cli.Command) error {

	envName := cmd.StringArg("name")
	if envName == "" {
		return cli.Exit("Please provide the name of the env environment to switch to", 1)
	}

	fmt.Printf("Switching to env file: [%s]\n", envName)

	// desiredFilePath := common.CustomEnvPrefix + cmd.Args().Get(0)

	// if os.IsNotExist(os.OpenFile() desiredFilePath) {
	// 	return cli.Exit("The specified .env file does not exist", 1)
	// }

	// if err := os.(desiredFilePath, common.BaseEnvFile); err != nil {

	return nil
}
