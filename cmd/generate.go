package cmd

import (
	"fmt"
	"io"
	"os"

	"github.com/SUSE/helm-certgen/pkg/certgen"
	"github.com/spf13/cobra"
)

type generateCmd struct {
	namespace string
	certyaml  string
	out       io.Writer
}

const generateDesc = `
This command creates certificates as per the configuration specified in certs.yaml.`

func newGenerateCmd(out io.Writer) *cobra.Command {
	gc := &generateCmd{
		out: out,
	}
	cmd := &cobra.Command{
		Use:          "generate [CHART]",
		Short:        "generate certificate",
		Long:         generateDesc,
		SilenceUsage: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) <= 0 {
				fmt.Printf("Please specify the chart to be installed.")
				fmt.Printf("\nGet help on how to use generate command with, helm certgen generate --help")

				return nil
			}
			return gc.run(args)
		},
	}
	f := cmd.Flags()
	f.StringVarP(&gc.namespace, "namespace", "n", "", "namespace to install the release into")
	f.StringVarP(&gc.certyaml, "certyaml", "c", "", "specify certs in a YAML file")
	return cmd
}

func (g *generateCmd) run(args []string) error {
	var certyaml string
	if len(g.certyaml) == 0 {
		certyaml = args[0] + "/certs.yaml"
		if _, err := os.Stat(certyaml); os.IsNotExist(err) {
			return fmt.Errorf("certyaml file not found in the specified chart")
		}
	} else {
		certyaml = g.certyaml
		if _, err := os.Stat(certyaml); os.IsNotExist(err) {
			return fmt.Errorf("certyaml file %s not found", certyaml)
		}
	}
	certSecret, _ := certgen.GetCertConfig(certyaml)
	fmt.Printf("Creating certificates for %s", certSecret.Name)

	certGen := certgen.New(g.namespace)
	certGen.GenerateCertificate(certSecret)
	return nil
}
