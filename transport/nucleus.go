package transport

// Nucleus structure
type Nucleus struct {
	dna Dna
}

// Insert the dna into the nucleus
func (n *Nucleus) Insert(d *Dna) *Nucleus {
	n.dna = *d
	return n
}

// Extract dna
func (n *Nucleus) Extract() Dna {
	return n.dna
}
