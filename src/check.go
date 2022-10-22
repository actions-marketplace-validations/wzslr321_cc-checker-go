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

func CheckFiles() error {
	err := walkThroughFiles()
	if err != nil {
		log.Fatalf("Maximum complexity reached: %s", err)
	}

	return nil
}

func scanFile(filePath string) error {
	file, err := os.Open(filePath)
	if err != nil {
		// do sth
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
				cc, _ := checkFunctionComplexity(currentFunc)
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
		cc, _ := checkFunctionComplexity(currentFunc)
		if cc > cnf.MaxComplexity {
			return &ccError{filePath: filePath, line: line, ccScore: cc}
		}
	}

	return nil
}

func checkFunctionComplexity(function string) (int, error) {
	fmt.Printf("Checking function: %s\n", function)

	return 0, nil
}

func walkThroughFiles() error {
	if isFileExcluded() {
		return nil
	}

	err := scanFile("test.txt")
	if err != nil {
		log.Fatalf("Maximum complexity reached: \n%s", err)
	}

	return nil
}

func isFileExcluded() bool {
	_, err := regexp.Compile("foo")
	if err != nil {
		return false
	}

	return false
}
