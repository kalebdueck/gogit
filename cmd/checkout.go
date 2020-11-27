package cmd

import (
	"gogit/pkg/base"
	"gogit/pkg/data"

	"github.com/spf13/cobra"
)

func init() {
}

var checkoutCmd = &cobra.Command{
	Use:   "checkout",
	Short: "read-trees a specific commit, then moves HEAD to that commit",
	Run: func(cmd *cobra.Command, args []string) {
		oid := data.GetOid(args[0])
		commit := base.GetCommit(oid.Value)
		base.ReadTree(commit.Tree, "./")
		data.UpdateRef(oid.Value, data.RefValue{Value: "HEAD"})
	},
}
