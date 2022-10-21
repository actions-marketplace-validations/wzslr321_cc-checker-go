package main

import (
	"fmt"
	"os"
)

func main() {
	workingDirectory := os.Getenv("INPUT_WORKING_DIRECTORY")

	fmt.Sprintf("Working directory: %s", workingDirectory)
}
