package transport

// Dna struct
type Dna struct {
	code []string
}

// Create the dna
// take all the functions and zip them
// func Create() Dna {
// 	return Dna{}
// }


// Add code
func (d *Dna) Add(codes ...string) Dna {

	for _, code := range codes {
		d.code = append(d.code, code)
	}
	return *d
}

func (d *Dna) Extract() []string {
   return d.code
}