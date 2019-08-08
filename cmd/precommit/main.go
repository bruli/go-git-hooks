package main

import (
	"fmt"
	"log"
	"os/exec"
)

func main() {
	fmt.Println("Execution pre-commit hook...")
	cmd := "git"
	args := []string{"rev-parse", "--verify", "HEAD"}
	out, err := exec.Command(cmd, args...).Output()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s", out)
}
