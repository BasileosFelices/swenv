package commands

import (
	"context"
	"fmt"
	"io"
	"os"

	"github.com/BasileosFelices/swenv/internal/commands/common"
	"github.com/urfave/cli/v3"
)

func copyFile(src string, dst string) error {

	sourceRead, err := os.Open(src)
	if err != nil {
		return err
	}
	defer sourceRead.Close()

	destWrite, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer destWrite.Close()

	_, err = io.Copy(destWrite, sourceRead)
	return err
}

func SwitchEnvFile(ctx context.Context, cmd *cli.Command) error {

	envName := cmd.StringArg("name")
	if envName == "" {
		return cli.Exit("Please provide the name of the env environment to switch to", 1)
	}

	// backup current .env if exists
	if _, err := os.Stat(common.BaseEnvFile); err == nil {
		err = copyFile(common.BaseEnvFile, common.BackupEnvFile)
		if err != nil {
			return cli.Exit(fmt.Sprintf("Failed to backup current .env file: %v", err), 1)
		}
		fmt.Println("Backed up current .env to .env.backup")
	}

	srcFile := common.CustomEnvPrefix + envName // ".env.<env>"

	fmt.Printf("Switching to env file: [%s]\n", envName)

	err := copyFile(srcFile, common.BaseEnvFile)
	if err != nil {
		if os.IsNotExist(err) {
			return cli.Exit(fmt.Sprintf("Env file [%s] does not exist", srcFile), 1)
		}
		return cli.Exit(fmt.Sprintf("Failed to switch to env file [%s]: %v", envName, err), 1)
	}

	return nil
}
