package cmd

import (
	"fmt"

	//"github.com/Arjun-P-Jayakrishnan/lcvs-cli.git/utils"
	"github.com/spf13/cobra"
)


var initCmd = &cobra.Command{
	Use:   "init",
	Short: "initialize LCVS for current project",
	Run: func(cmd *cobra.Command, args []string) { 
		fmt.Println("initialised",args[0],args[1]);

		//projectPath,err:=utils.AbsPath(".")

		// if err!=nil{
		// 	fmt.Printf("Error getting absolute path %w",err)
		// }



	},
}


func init() {

	rootCmd.AddCommand(initCmd)
}


