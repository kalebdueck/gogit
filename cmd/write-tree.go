package cmd

import (
	"fmt"
	"gogit/pkg/base"

	"github.com/spf13/cobra"
)

var writeTreeCmd = &cobra.Command{
	Use:   "write-tree",
	Short: "initializes a gogit repository",
	Run: func(cmd *cobra.Command, args []string) {
		result := base.WriteTree(args[0])
		fmt.Println(result)
	},
}
