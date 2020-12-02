package cmd

import (
	"fmt"
	"gogit/pkg/base"

	"github.com/spf13/cobra"
)

var oidFlag string

var logCmd = &cobra.Command{
	Use:   "log",
	Short: "initializes a gogit repository",
	Run: func(cmd *cobra.Command, args []string) {

		var oid string
		if len(args) > 0 {
			oid = args[0]
		}
		oid = base.GetOid(oid)

		fmt.Println("logging")
		fmt.Println(oid)
		for _, commitOid := range base.IterCommitsAndParents([]string{oid}) {
			commit := base.GetCommit(commitOid)

			fmt.Printf("commit: %s\n", commitOid)
			fmt.Printf("message: %s\n", commit.Message)
			fmt.Printf("-----------\n")

		}
	},
}
