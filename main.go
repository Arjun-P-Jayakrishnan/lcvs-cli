/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"

	"github.com/Arjun-P-Jayakrishnan/lcvs-cli.git/api"
	"github.com/Arjun-P-Jayakrishnan/lcvs-cli.git/cmd"
	//"github.com/Arjun-P-Jayakrishnan/lcvs-cli.git/api"
)

func main() {
	/*
		Load environment variables
	*/
	envErr := godotenv.Load(".env")
	if envErr!=nil{
		log.Fatal(envErr)
	}

	portString := os.Getenv("PORT")
	if portString == "" {
		log.Fatal("No Port detected")
	}

	routerDB := api.Router {
		Port: portString,
	}

	err :=routerDB.Init()
	if err!=nil {
		log.Fatal(err)
	}

	cmd.Execute()
}


