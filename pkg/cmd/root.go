package cmd

import (
	"fmt"
	"os"

	"github.com/cnoe-io/idpbuilder/pkg/cmd/create"
	"github.com/cnoe-io/idpbuilder/pkg/cmd/gitserver"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "idpbuilder",
	Short: "Manage reference IDPs",
	Long:  "",
}

func init() {
	rootCmd.AddCommand(create.CreateCmd)
	rootCmd.AddCommand(gitserver.GitServerCmd)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
