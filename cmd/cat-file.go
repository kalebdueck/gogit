package cmd

import (
	"bufio"
	"gogit/pkg/base"
	"gogit/pkg/data"
	"os"

	"github.com/spf13/cobra"
)

func init() {
	catFileCmd.Flags().StringVar(&expectedFlag, "expected", "blob", "Expected type of Object")
}

var expectedFlag string

var catFileCmd = &cobra.Command{
	Use:   "cat-file",
	Short: "Echos a file to stdout",
	Run: func(cmd *cobra.Command, args []string) {

		oid := base.GetOid(args[0])
		f := bufio.NewWriter(os.Stdout)
		defer f.Flush()
		resp, err := data.GetObject(oid, expectedFlag)

		if err != nil {
			panic(err)
		}

		f.Write(resp)
	},
}
