package checker

import (
	"github.com/wzslr321/cc-checker-go/config"
	"log"
	"regexp"
	"time"
)

var cnf = config.GetConfig()

func CheckFiles() error {
	err := walkThroughFiles()
	if err != nil {
		log.Fatalf("Maximum complexity reached: %s", err)
	}

	return nil
}

func checkFileComplexity() error {
	if time.Now().Hour() < cnf.MaxComplexity {
		return nil
	}

	return nil
}

func walkThroughFiles() error {
	if isFileExcluded() {
		return nil
	}

	err := checkFileComplexity()
	if err != nil {
		return err
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
