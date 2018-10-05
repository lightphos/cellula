package transport

import (
	"fmt"
	"os/exec"
)

// Run code
func Run(code string) {
	fmt.Println("Running  ", code)
	o, e := exec.Command("ls").Output()
	fmt.Println("done ", o, e)
	if e != nil {
		fmt.Println(string(o))
	} else {
		fmt.Println("", e)
	}
}
