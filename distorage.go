package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"strings"
)

func main() {

	var cmdStore = &cobra.Command{
		Use:   "store [filename]",
		Short: "store creates encrypted file blocks",
		Long: `store is used to split and encrypt files for distribute
	storage.
        `,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Print: " + strings.Join(args, " "))
		},
	}

	var cmdVersion = &cobra.Command{
		Use:   "version",
		Short: "Print the version number of Distorage",
		Long:  `All software has versions. This is Distorage's`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Distorage distributed file storage v0.1 -- HEAD")
		},
	}

	var rootCmd = &cobra.Command{Use: "distorage"}
	rootCmd.AddCommand(cmdStore, cmdVersion)
	rootCmd.Execute()
}
