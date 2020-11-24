package cmd

import (
	"fmt"
	"gogit/pkg/base"
	"gogit/pkg/data"

	"github.com/spf13/cobra"
)

var kCmd = &cobra.Command{
	Use:   "k",
	Short: "Prints Refs",
	Run: func(cmd *cobra.Command, args []string) {
		var oids []string
		for refname, ref := range data.IterRefs() {
			fmt.Println(fmt.Sprintf("%s: %s", refname, ref))
			oids = append(oids, ref)
		}

		for _, oid := range base.IterCommitsAndParents(oids) {
			commit := base.GetCommit(oid)
			fmt.Println(oid)
			if commit.Parent != "" {
				fmt.Printf("Parent: %s \n", commit.Parent)

			}

		}
	},
}
