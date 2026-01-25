package main

import (
	"context"
	"log"
	"os"

	"github.com/BasileosFelices/swenv/internal/commands"
)

func main() {
	cmd := commands.NewRootCommand()
	if err := cmd.Run(context.TODO(), os.Args); err != nil {
		log.Fatal(err)
	}
}
