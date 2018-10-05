package transport

import (
	"fmt"
)

// Cell struct
type Cell struct {
	n Nucleus
	UUID string
}

// CreateCell the cell
// take all the functions and zip them
func CreateCell() Cell {
	var c = Cell{UUID: "1212"}
	return c
}

// Insert the nucleus into the cell
func (c *Cell) Insert(n *Nucleus) Cell {
	c.n = *n
	return *c
}

func (c *Cell) Start() {
  nu := c.n
  d := nu.Extract()
  for _, cd := range d.Extract() {
	fmt.Print(" ", cd)
  }

	fmt.Println()
}