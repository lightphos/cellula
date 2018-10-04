package cell

import (
	nucleus "cellula/nucleus"
)

// Cell struct
type Cell struct {
	n    nucleus.Nucleus
	UUID string
}

// Create the cell
// take all the functions and zip them
func Create() Cell {
	var c = Cell{UUID: "1212"}
	return c
}

// Insert the nucleus into the cell
func (c *Cell) Insert(n *nucleus.Nucleus) Cell {
	c.n = *n
	return *c
}
