package cmd

import (
	"fmt"
	"gogit/pkg/base"
	"gogit/pkg/data"

	"github.com/spf13/cobra"
)

var tagCmd = &cobra.Command{
	Use:   "tag",
	Short: "provides a mechanism for naming commits",
	Args:  cobra.RangeArgs(1, 2),
	Run: func(cmd *cobra.Command, args []string) {
		tagName := args[0]
		fmt.Println(tagName)

		//Todo cobra default args?
		var oid string
		if len(args) > 1 {
			oid = args[1]
		} else {
			oid = data.GetRef("HEAD")
		}

		base.CreateTag(tagName, oid)
	},
}
