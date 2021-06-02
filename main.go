package main

import (
	"fmt"
	"os"

	"github.com/SoloDeploy/solo-providers-git-github/cmd"
)

func main() {
	command := cmd.NewCmd()

	if err := command.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
