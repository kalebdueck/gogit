package cmd

import (
	"fmt"
	"gogit/pkg/base"

	"github.com/spf13/cobra"
)

func init() {
	commitCmd.Flags().StringVarP(&messageFlag, "message", "m", "", "Commit message")
	commitCmd.MarkFlagRequired("message")
}

var messageFlag string

var commitCmd = &cobra.Command{
	Use:   "commit",
	Short: "initializes a gogit repository",
	Run: func(cmd *cobra.Command, args []string) {
		result := base.Commit(messageFlag)

		fmt.Println(result)
	},
}
