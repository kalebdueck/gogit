package cmd

import (
	"gogit/pkg/base"

	"github.com/spf13/cobra"
)

var readTreeCmd = &cobra.Command{
	Use:   "read-tree",
	Short: "initializes a gogit repository",
	Run: func(cmd *cobra.Command, args []string) {
		oid := base.GetOid(args[0])
		base.ReadTree(oid, "./")
	},
}
