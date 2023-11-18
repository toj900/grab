package cmd

import (
	"os"

	"k8s.io/cli-runtime/pkg/genericiooptions"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/toj900/grab/pkg/schema"
)

var rootCmd = schema.NewCmdSchema(genericiooptions.IOStreams{In: os.Stdin, Out: os.Stdout, ErrOut: os.Stderr})

func Execute() error {
	flags := pflag.NewFlagSet("kubectl-grabschema", pflag.ExitOnError)
	rootCmd.CompletionOptions = cobra.CompletionOptions{DisableDefaultCmd: true}
	pflag.CommandLine = flags
	if err := rootCmd.Execute(); err != nil {
		return err
	}
	return nil
}
