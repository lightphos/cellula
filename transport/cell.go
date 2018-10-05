package transport

import (
	"fmt"
)

// Cell struct
type Cell struct {
	n    Nucleus
	UUID string
}

// CreateCell the cell
// take all the functions and zip them
func CreateCell() Cell {
	var c = Cell{UUID: "1212"}
	return c
}

// Insert the nucleus into the cell
func (c *Cell) Insert(n *Nucleus) *Cell {
	c.n = *n
	return c
}

// Start cell
// Ok what shall we do now....
// First cell -> nucleus -> dna -> code
func (c *Cell) Start() {
	nucleus := c.n
	dna := nucleus.Extract()
	for _, c := range dna.Extract() {
		fmt.Print("run ", c)
		Run(c)
	}

	fmt.Println()
}
