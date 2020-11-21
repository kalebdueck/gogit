package cmd

import (
	"bufio"
	"fmt"
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
	Short: "initializes a gogit repository",
	Run: func(cmd *cobra.Command, args []string) {

		object := base.GetOid(args[0])

		expected := []byte(object)
		f := bufio.NewWriter(os.Stdout)
		defer f.Flush()
		resp, err := data.GetObject(object, expected)

		if err != nil {
			fmt.Println(err)
			return
		}

		f.Write(resp)
	},
}
