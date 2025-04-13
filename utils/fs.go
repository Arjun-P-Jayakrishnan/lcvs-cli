package utils

import (
	"bufio"
	"io/fs"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

// Dir
func CreateDir(path string) error {
	return os.Mkdir(path, 0755)
}

func DirOrFileExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

func RemoveDir(path string) error {
	err := os.RemoveAll(path)

	return err
}

func GetCWD() (string, error) {
	return os.Getwd()
}

// Path
func PathJoin(parts ...string) string {
	return filepath.Join(parts...)
}

func AbsPath(path string) (string, error) {
	return filepath.Abs(path)
}

// Plain Text
func ReadFile(path string) ([]byte, error) {

	return ioutil.ReadFile(path)
}

func WriteFile(path string, data []byte) error {

	return ioutil.WriteFile(path, data, 0755)
}

// Parsing
func ReadLines(path string) ([]string, error) {
	file, err := os.Open(path)

	if err != nil {
		return nil, err
	}

	defer file.Close()

	var lines []string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines, scanner.Err()
}

func WalkFiles(root string) ([]string, error) {
	var files []string

	err := filepath.Walk(root, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			//Skip this path but continue walking
			return nil
		}

		//Skips hidden files or folders
		base := filepath.Base(path)
		if strings.HasPrefix(base, ".") {
			if path != root {

				if info.IsDir() {
					return filepath.SkipDir
				}

				return nil
			}
		}

		//Skips symlinks - basically a file that behaves like a pointer to external files
		if info.Mode()&os.ModeSymlink != 0 {
			return nil
		}

		//Add only regular files
		if !info.IsDir() {
			files = append(files, path)
		}

		return nil
	})

	return files, err
}


func AppendToGitIgnore(root string) error{
	gitIgnorePath := filepath.Join(root,".gitignore")
	entries:=[]string{".lcvs/",".lcvs-ignore"}

	existing :=map[string]bool{}
	if DirOrFileExists(gitIgnorePath){
		lines,err := ReadLines(gitIgnorePath)
		if err!=nil{
			return err
		}

		for _,line :=range lines{
			existing[strings.TrimSpace(line)]=true
		}
	}

	file,err:=os.OpenFile(gitIgnorePath,os.O_APPEND|os.O_CREATE|os.O_WRONLY,0644)
	if err!=nil{
		return err
	}

	defer file.Close()

	writer := bufio.NewWriter(file)

	for _,entry :=range entries {
		if !existing[entry] {
			if _,err := writer.WriteString("\n"+entry+"\n");err!=nil{
				return err
			}
		}
	}

	return writer.Flush()
}