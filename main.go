/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"fmt"
	"log"

	"github.com/Arjun-P-Jayakrishnan/lcvs-cli.git/utils"
)

//"github.com/joho/godotenv"

//"github.com/Arjun-P-Jayakrishnan/lcvs-cli.git/api"
//"github.com/Arjun-P-Jayakrishnan/lcvs-cli.git/cmd"
//"github.com/Arjun-P-Jayakrishnan/lcvs-cli.git/api"

func main() {
	
	//cmd.Execute()

	tmpPath:="./temp-Test"

	if err:=utils.CreateDir(tmpPath);err!=nil{
		log.Fatal("Cant create directory",err)
	}
	fmt.Printf("Directory Created")

	textPath:=utils.PathJoin(tmpPath,"text.txt")
	if err:=utils.WriteFile(textPath,[]byte("hello from lcvs cli"));err!=nil{
		log.Fatal("Cant write to text file",err)
	}
	fmt.Println("Wrote Plain text file")

	content,err:=utils.ReadFile(textPath)
	if err!=nil{
		log.Fatal("Error Reading from file",err)
	}

	fmt.Print(string(content))

	meta:= map[string] string{
		"name" :"Sample",
		"owner":" author",
	}

	jsonPath := utils.PathJoin(tmpPath,"meta.json")

	if err := utils.WriteJSON(jsonPath,meta);err!=nil{
		panic("WriteJSON failed :"+err.Error())
	}
	fmt.Printf("Write json")

	var loaded map[string]string
	if err:=utils.ReadJSON(jsonPath,&loaded);err!=nil{
		panic("Error rading json:"+err.Error())
	}
	fmt.Println(loaded)

	abs,err:=utils.AbsPath(textPath)

	if err!=nil{
		panic("Error in getting abs path "+err.Error())
	}

	fmt.Print(abs)
}


