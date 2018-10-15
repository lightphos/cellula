package blueprint

import (
	"context"
	"fmt"
	"log"
	"os/exec"
	"strings"

	"docker.io/go-docker/api/types"
)

// Run me
func Run() {
	fns, err := exec.Command("ls").Output()
	if err != nil {
		log.Fatal(":--->", err)
	}

	s := string(fns)
	slice := strings.Split(s, "\n")
	for _, fn := range slice {
		fmt.Println("--> " + fn)
	}

	// output, err := exec.Command("go", "run", "../cellula.go").Output()
	// if err != nil {
	// 	log.Fatal("--->", err)
	// }
	// fmt.Printf("%s\n", output)

	cli, err := docker.NewEnvClient()
	if err != nil {
		panic(err)
	}

	containers, err := cli.ContainerList(context.Background(), types.ContainerListOptions{})
	if err != nil {
		panic(err)
	}

	for _, container := range containers {
		fmt.Printf("%s %s\n", container.ID[:10], container.Image)
	}
}
