package cmd

import (
	"fmt"
	"gogit/pkg/data"
	"io/ioutil"

	"github.com/spf13/cobra"
)

func init() {
	hashObjectCmd.Flags().StringVar(&typeFlag, "type", "blob", "Type of Object")
}

var typeFlag string

var hashObjectCmd = &cobra.Command{
	Use:   "hash-object",
	Short: "",
	Run: func(cmd *cobra.Command, args []string) {
		file := args[0]
		types := []byte(typeFlag)
		dat, err := ioutil.ReadFile(file)
		check(err)
		fmt.Println(data.HashObject(dat, types))
	},
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
