package cmd

import (
	"fmt"
	"gogit/pkg/base"
	"gogit/pkg/data"
	"strings"

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

			var newOid string = ""
			for _, line := range strings.Split(commit, "\n") {
				split := strings.Split(line, " ")
				if split[0] == "parent" {
					newOid = split[1]
				}
			}

			oid = newOid
		}
	},
}
