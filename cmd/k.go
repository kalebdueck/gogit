package cmd

import (
	"fmt"
	"gogit/pkg/data"

	"github.com/spf13/cobra"
)

var kCmd = &cobra.Command{
	Use:   "k",
	Short: "Prints Refs",
	Run: func(cmd *cobra.Command, args []string) {
		for refname, ref := range data.IterRefs() {
			fmt.Println(fmt.Sprintf("%s: %s", refname, ref))
		}
	},
}
