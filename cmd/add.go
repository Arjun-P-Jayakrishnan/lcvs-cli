/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/Arjun-P-Jayakrishnan/lcvs-cli.git/internal"
	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add files[...]",
	Short: "add files to the staging area",
	Long: `Stage one or more files for the next commit.`,
	Run: func(cmd *cobra.Command, args []string) {
		
		if len(args)==0{
			args = []string{"."}
		}
		//Core implementation
		err := internal.AddFiles(args)

		if err!=nil {
			fmt.Println("Error staging file :",err)
			os.Exit(1)
		}

		fmt.Println("✅ Files Staged successfully.")
	},
}

func init() {
	rootCmd.AddCommand(addCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
