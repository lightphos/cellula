package nucleus

import "cellula/dna"

// Nucleus structure
type Nucleus struct {
	dna dna.Dna
}

// Create the nucleus of code
// take all the functions and zip them
func Create() Nucleus {
	return Nucleus{}
}

// Insert the dna into the nucleus
func (n *Nucleus) Insert(d *dna.Dna) Nucleus {
	n.dna = *d
	return *n
}
