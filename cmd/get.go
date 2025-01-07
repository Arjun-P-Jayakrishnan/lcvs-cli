/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"os"
	"io"
	"net/http"
	"github.com/spf13/cobra"
)

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		

		var goferName = ""

		if len(args) > 1 && args[0] != ""{
			goferName=args[0]
		}

		URL:=""

		fmt.Println("Try to get"+goferName+" gopher")

		//
		resposne,err:=http.Get(URL)

		if err!=nil{
			fmt.Println(err)
		}

		defer resposne.Body.Close()

		if resposne.StatusCode == 200 {
			out,err :=os.Create(goferName+".png")
			if err!=nil {
				fmt.Println(err)
			}

			defer out.Close()

			_,err=io.Copy(out,resposne.Body)
			if err !=nil {
				fmt.Println(err)
			}

			fmt.Println("Perfect! Just Svaed in"+out.Name()+"!")
		}else{
			fmt.Println("Error:"+goferName+" does  not exist")
		}
	},
}

func init() {
	rootCmd.AddCommand(getCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
