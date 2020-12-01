package cmd

import (
	"fmt"
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
		name := args[0]
		oid := data.GetOid(name)
		commit := base.GetCommit(oid.Value)
		base.ReadTree(commit.Tree, "./")

		var head data.RefValue
		if base.IsBranch(name) {
			head = data.RefValue{
				Value:    fmt.Sprintf("refs/heads/%s", name),
				Symbolic: true,
			}
		} else {
			head = data.RefValue{
				Value:    oid.Value,
				Symbolic: false,
			}
		}

		data.UpdateRef("HEAD", head, false)
	},
}
