package src

import (
	"bufio"
	"fmt"
	"github.com/wzslr321/cc-checker-go/config"
	"log"
	"os"
	"regexp"
)

var cnf = config.GetConfig()

type ccError struct {
	line     int
	ccScore  int
	filePath string
}

func (e *ccError) Error() string {
	return fmt.Sprintf("Maximum complexity reached in file %s at line %d\n CC score: %d", e.filePath, e.line, e.ccScore)
}

func CheckFiles(dir string) error {
	if dir == "" {
		dir = cnf.WorkDir
	}

	files, err := os.ReadDir(dir)
	if err != nil {
		return err
	}

	for _, file := range files {
		filePath := dir + "/" + file.Name()
		if !file.IsDir() {
			if isFileExcluded(file.Name()) {
				continue
			}
			err := scanFile(filePath)
			if err != nil {
				log.Fatalf("Maximum complexity reached: \n%s", err)
			}
		} else {
			err := CheckFiles(filePath)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func scanFile(filePath string) error {
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	line := 0
	scanningFunc := false
	currentFunc := ""
	for scanner.Scan() {
		lineStr := scanner.Text()
		if scanningFunc && len(lineStr) > 4 {
			if lineStr[0:4] == "func" {
				scanningFunc = false
				cc, _ := getFunctionComplexity(currentFunc)
				if cc > cnf.MaxComplexity {
					return &ccError{filePath: filePath, line: line, ccScore: cc}
				}
				currentFunc = ""
			} else {
				currentFunc += lineStr
			}
		}
		if !scanningFunc && len(lineStr) > 4 {
			if lineStr[0:4] == "func" {
				scanningFunc = true
				currentFunc += lineStr
			}
		}
		line++
	}
	if len(currentFunc) > 0 {
		cc, _ := getFunctionComplexity(currentFunc)
		if cc > cnf.MaxComplexity {
			return &ccError{filePath: filePath, line: line, ccScore: cc}
		}
	}

	return nil
}

func getFunctionComplexity(function string) (int, error) {
	fmt.Printf("Checking function: %s\n", function)
	// it should probably contain some real functionality 🌚

	return 0, nil
}

func isFileExcluded(filename string) bool {
	for _, pattern := range cnf.ExcludePatterns {
		match, _ := regexp.MatchString(pattern, filename)
		if match {
			return true
		}
	}

	return false
}
