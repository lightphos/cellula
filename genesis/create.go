package main

import (
	cell "cellula/cell"
	dna "cellula/dna"
	nucleus "cellula/nucleus"
	"fmt"
)

func main() {
	d := dna.Create()

	n := nucleus.Create()
	n.Insert(&d)
	c := cell.Create()
	c.Insert(&n)
	fmt.Println("genesis", c)
}
