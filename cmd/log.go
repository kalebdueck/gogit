package cmd

import (
	"fmt"
	"gogit/pkg/base"
	"gogit/pkg/data"

	"github.com/spf13/cobra"
)

func init() {
	logCmd.Flags().StringVar(&oidFlag, "oid", "", "Expected type of Object")
}

var oidFlag string

var logCmd = &cobra.Command{
	Use:   "log",
	Short: "initializes a gogit repository",
	Run: func(cmd *cobra.Command, args []string) {
		oid := oidFlag
		if oid == "" {
			oid = data.GetHead()
		}

		for oid != "" {
			commit := base.GetCommit(oid)

			fmt.Printf("commit: %s\n", oid)
			fmt.Printf("message: %s\n", commit.Message)
			fmt.Printf("-----------\n")

			oid = commit.Parent
		}
	},
}
