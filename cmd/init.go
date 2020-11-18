package cmd

import (
	"fmt"
	"gogit/pkg/data"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(initCmd)
}

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "initializes a gogit repository",
	Run: func(cmd *cobra.Command, args []string) {
		data.Init()
		fmt.Println(".gogit repository initialized")
	},
}
