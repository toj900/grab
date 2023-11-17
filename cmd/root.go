package cmd

import (
	"github.com/spf13/cobra"

	"github.com/toj900/grab/pkg/schema"
)

const (
	grabSchemaLongDescription = `
Create example CustomResources from CRDs
 $ grab-schema
`
	grabSchemaExamples = `
  Create example CustomResources from CRDs
   $ grab-schema
`
)

var rootCmd = &cobra.Command{
	Use:     "grab-schema",
	Short:   "Create example YAMLs from CRDs",
	Long:    grabSchemaLongDescription,
	Args:    cobra.NoArgs,
	Example: grabSchemaExamples,
	Run: func(cmd *cobra.Command, args []string) {
		schema.GrabCRD(resources)
	},
}
var resources []string
var resources1 []string

func init() {
	rootCmd.Flags().StringArrayVarP(&resources, "resource", "r", resources, "ussage")

}
func Execute() error {
	// rootCmd.SetOut(ketallOptions.Streams.Out)
	// rootCmd.SetErr(ketallOptions.Streams.ErrOut)
	return rootCmd.Execute()
}
