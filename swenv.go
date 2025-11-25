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
	}

	if err := cmd.Run(context.TODO(), os.Args); err != nil {
		log.Fatal(err)
	}
}
