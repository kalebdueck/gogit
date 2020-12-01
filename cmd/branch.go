package cmd

import (
	"fmt"
	"gogit/pkg/base"

	"github.com/spf13/cobra"
)

var branchCmd = &cobra.Command{
	Use:   "branch",
	Short: "you know what it does",
	Args:  cobra.RangeArgs(1, 2),
	Run: func(cmd *cobra.Command, args []string) {
		branchName := args[0]
		startPoint := "HEAD"
		if len(args) > 1 {
			startPoint = args[1]
		}

		base.CreateBranch(branchName, startPoint)

		startAbbrev := startPoint
		if len(startPoint) > 10 {
			startAbbrev = startPoint[:10]
		}

		fmt.Printf("Branch %s created at %s\n", branchName, startAbbrev)
	},
}
