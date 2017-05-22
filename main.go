package main

import (
	"os"

	"github.com/SUSE/helm-certgen/cmd"
)

func main() {
	cmd := cmd.NewRootCmd(os.Stdout)
	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}
