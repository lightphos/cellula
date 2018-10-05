package transport

import info "cellula/info"

// Dna struct
type Dna struct {
	code []info.Coder
}

// Create the dna
// take all the functions and zip them
// func Create() Dna {
// 	return Dna{}
// }

// Add code
func (d *Dna) Add(codes ...info.Coder) *Dna {

	for _, code := range codes {
		d.code = append(d.code, code)
	}
	return d
}

// Extract code
func (d *Dna) Extract() []info.Coder {
	return d.code
}
