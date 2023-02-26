package main

import (
	"fmt"
	"log"
	"os/exec"
)

func main() {
	// Define `grep` command.
	grep := exec.Command("grep", "-r", "err", ".")
	out, err := grep.StdoutPipe()
	if err != nil {
		log.Fatal(err)
	}

	// Start grep.
	if err := grep.Start(); err != nil {
		log.Fatal(err)
	}

	// Define `wc` command.
	wc := exec.Command("wc", "-l")
	wc.Stdin = out

	// Read combined output.
	data, err := wc.CombinedOutput()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s", data)
}
