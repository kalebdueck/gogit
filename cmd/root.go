package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "gogit",
	Short: "A simplified Go implementation of Git",
	Long:  "",
}

func init() {
	rootCmd.AddCommand(initCmd)
	rootCmd.AddCommand(hashObjectCmd)
	rootCmd.AddCommand(catFileCmd)
	rootCmd.AddCommand(writeTreeCmd)
	rootCmd.AddCommand(readTreeCmd)
	rootCmd.AddCommand(logCmd)
	rootCmd.AddCommand(commitCmd)
	rootCmd.AddCommand(checkoutCmd)
	rootCmd.AddCommand(tagCmd)
}

//Execute runs the commands
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
