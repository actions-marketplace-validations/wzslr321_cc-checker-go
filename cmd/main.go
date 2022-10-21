package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func listFiles() ([]string, error) {
	workingDirectory := os.Getenv("INPUT_WORKING_DIRECTORY")

	fileList := []string{}
	err := filepath.Walk(fmt.Sprintf("./%s", workingDirectory), func(path string, f os.FileInfo, err error) error {
		if doesFileMatch(path) {
			fileList = append(fileList, path)
		}
		return nil
	})
	return fileList, err
}

func doesFileMatch(path string) bool {
	if file, err := os.Stat(path); err == nil && !file.IsDir() {
		return true
	}
	return false
}

func main() {
	_, err := listFiles()
	if err != nil {
		log.Fatal("Fatalny błąd...")
	}
}
