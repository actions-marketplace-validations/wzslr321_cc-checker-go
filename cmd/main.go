package main

import (
	checker "github.com/wzslr321/cc-checker-go/src"
	"log"
)

func main() {
	err := checker.CheckFiles("")
	if err != nil {
		log.Fatalf("Error: %s", err)
	}
}
