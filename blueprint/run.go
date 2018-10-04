package main

import (
	"fmt"
	"log"
	"os/exec"
	"strings"
)

func main() {
	fns, err := exec.Command("ls", "../functions/*.go").Output()
	if err != nil {
		log.Fatal("--->", err)
	}

	s := string(fns)
	slice := strings.Split(s, "\n")
	for _, fn := range slice {
		fmt.Println("--> " + fn)
	}

	output, err := exec.Command("go", "run", "../transport/transport.go").Output()
	if err != nil {
		log.Fatal("--->", err)
	}
	fmt.Printf("%s\n", output)
}
