/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "lcvs",
	Short: "Local Version Control System for semantic versioning",
	Long: `LCVS or Local Version Control System is a version control documentation tool
	 that focuses on research oriented version control and safe stores.
		
	The main focus of this app is to create a system for sematic versioning of your save files
	such that it gives you access to code in form of a research module
	`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
