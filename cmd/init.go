package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)


var initCmd = &cobra.Command{
	Use:   "init",
	Short: "initialize repository",
	Long: `It initializes the save system.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) { 
		fmt.Println("initialised",args[0],args[1]);
	},
}


func init() {

	rootCmd.AddCommand(initCmd)
}


