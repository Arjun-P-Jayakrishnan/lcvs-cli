package cmd

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/Arjun-P-Jayakrishnan/lcvs-cli.git/internal"
	"github.com/Arjun-P-Jayakrishnan/lcvs-cli.git/utils"
	"github.com/spf13/cobra"
)

var (
	forceInit bool
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize LCVS Tracking for current project",
	Run: func(cmd *cobra.Command, args []string) {

		//Check for folder
		metaDir := filepath.Join(".", internal.MetaDirName)

		if utils.DirOrFileExists(metaDir) {

			if !forceInit {
				fmt.Println("⚠️\t Project already exists .Use --force to re intialize")
				return
			}

			fmt.Println(" ⚠️ Reinitializing,Deleting the .lcvs folder ")
			err := utils.RemoveDir(metaDir)

			if err != nil {
				fmt.Println("Error while removing the folder ", err)
				return
			}

		}

		//Get details for initailization
		reader := bufio.NewReader(os.Stdin)

		fmt.Print("Project Name : ")
		name, nameErr := reader.ReadString('\n')

		if nameErr != nil {
			log.Fatal("Project name error", nameErr)
			return
		}

		fmt.Print("Author : ")
		author, authErr := reader.ReadString('\n')

		if authErr != nil {
			log.Fatal("Author error", authErr)
			return
		}

		fmt.Print("Decription :")
		desc, descErr := reader.ReadString('\n')

		if descErr != nil {
			log.Fatal("Decsription Error ", descErr)
			return
		}

		dirErr := utils.CreateDir(metaDir)

		if dirErr != nil {
			fmt.Println("❌ Cant create directory")
			log.Fatal("❌ Can't Create Folder", dirErr)
			return
		}

		//Create file to store data
		metaPath := filepath.Join(metaDir, internal.MetaFileName)

		cwd, errWD := utils.GetCWD()
		if errWD != nil {
			fmt.Println("Cant get the current working directory")
			log.Fatal("Cant get current working directory :", errWD)
		}

		//Create and stor data in meta.json file
		projectMeta := internal.ProjectMeta{
			Name:          strings.TrimSpace(name),
			Author:        strings.TrimSpace(author),
			Description:   strings.TrimSpace(desc),
			CreatedAt:     time.Now().Format(time.RFC3339),
			ProjectID:     utils.GenerateUUID(),
			LastKnownPath: cwd,
		}

		err := utils.WriteJSON(metaPath, projectMeta)

		_=utils.AppendToGitIgnore(cwd)

		if err != nil {
			fmt.Println("❌ Error Creating JSON meta data file")
			log.Fatal("❌ Error Creating JSON meta data file")
			return
		}
	},
}

func init() {
	initCmd.Flags().BoolVarP(&forceInit, "force", "f", false, "force reinitialization (delete .lcvs)")
	rootCmd.AddCommand(initCmd)

}
