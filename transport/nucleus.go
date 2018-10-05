package transport

// Nucleus structure
type Nucleus struct {
	dna Dna
}

// Create the nucleus of code
// take all the functions and zip them
// func Create() Nucleus {
// 	return Nucleus{}
// }

// Insert the dna into the nucleus
func (n *Nucleus) Insert(d *Dna) *Nucleus {
	n.dna = *d
	return n
}

// Extract dna
func (n *Nucleus) Extract() Dna {
	return n.dna
}
