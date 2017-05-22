package cmd

import (
	"io"

	"github.com/spf13/cobra"
)

var globalUsage = `Plugin for Kubernetes helm CLI.
This supports 
	$ helm certgen generate`

func NewRootCmd(out io.Writer) *cobra.Command {

	cmd := &cobra.Command{
		Use:          "certgen",
		Short:        "certificate generator plugin for helm cli",
		Long:         globalUsage,
		SilenceUsage: true,
	}

	cmd.AddCommand(
		newGenerateCmd(out),
	)
	return cmd
}
