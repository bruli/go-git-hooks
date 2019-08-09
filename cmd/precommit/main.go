package main

import (
	"fmt"
	"github.com/bruli/go-git-hooks/git"
	"log"
)

func main() {
	fmt.Println("Executing pre-commit hook...")
	fil, err := git.GetCommittedFiles()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s", fil)
}
