package internal

import (
	"fmt"
	"log"
	"path/filepath"
	"strings"

	"github.com/Arjun-P-Jayakrishnan/lcvs-cli.git/utils"
)

// buildIgnoreSet builds the ignore set which can be looked up for ignored files and directories
//
//	`ignoreLines` []string lines inside the .lcvs-ignore file
func buildIgnoreSet(ignoreLines []string) map[string]bool {
	ignored := map[string]bool{}

	for _, line := range ignoreLines {
		line = strings.TrimSpace(line)
		//if line is not empty nor is a comment add it to list
		if line != "" && !strings.HasPrefix(line, "#") {
			ignored[line] = true
		}
	}

	return ignored
}

// isIgnored  checks if the file path is ignored or can be added
//
//	`file` 		string 		path to file
//	`ignoreSet` map[string]bool 	list of ignored paths
//	`root` 		string 		root path
func isIgnored(file string, ignoreSet map[string]bool, root string) bool {
	absPath, err := utils.AbsPath(file)

	if err != nil {
		log.Fatal("absolute path error", err)
		return false
	}

	//gets the relative path i.e removes the root string from files absolute path
	relPath, err := filepath.Rel(root, absPath)
	if err != nil {
		log.Fatal("relative path error ", err)
		return false
	}

	//we split the path
	parts := strings.Split(relPath, string(filepath.Separator))

	for i := 1; i <= len(parts); i++ {
		// we add one by one
		partial := filepath.Join(parts[:i]...)

		//if at any point it is added in the ignore set we ignore it
		if ignoreSet[partial] {
			return true
		}
	}

	return false
}

func AddFiles(paths []string) error {

	//check if .lcvs directory exists
	if !utils.DirOrFileExists(MetaDirName) {
		return fmt.Errorf("cant add files to an un initialized project")
	}

	rootDir, err := utils.GetCWD()
	if err != nil {
		return err
	}

	//Load ignore-list from .lcvs-ignore
	ignoreSet := map[string]bool{}

	//Build the Ignore Set
	if utils.DirOrFileExists(IgnoreFileName) {
		lines, err := utils.ReadLines(IgnoreFileName)

		if err == nil {
			ignoreSet = buildIgnoreSet(lines)
		}
	}

	//Collect valid files
	var filesToStage []StagedFile
	for _, input := range paths {

		files, err := utils.WalkFiles(input)

		if err != nil {
			return err
		}

		for _, file := range files {

			//If in ignore list ignore
			if isIgnored(file, ignoreSet, rootDir) {
				continue
			}

			hash, err := utils.HashFileSHA256(file)
			if err != nil {
				return err
			}

			//else add it to teh existing list
			filesToStage = append(filesToStage,
				StagedFile{
					Path: file,
					Hash: hash,
				})
		}
	}

	//Load existing staging data
	var stagingData StagingInfo
	if utils.DirOrFileExists(
		StagingFilePath) {
		_ = utils.ReadJSON(
			StagingFilePath, &stagingData)
	}

	//DeDuplicate by having
	// existing list
	// new files list
	//comapre and store in staging
	existing := map[string]bool{}
	for _, file := range stagingData.Files {
		existing[file.Path] = true
	}

	for _, file := range filesToStage {
		if !existing[file.Path] {
			stagingData.Files = append(stagingData.Files, file)
		}
	}

	return utils.WriteJSON(
		filepath.Join(MetaDirName, StagingFilePath), stagingData)
}
