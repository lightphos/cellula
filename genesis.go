package main

import (
	"fmt"
	t "cellula/transport"
)

func main() {
	d := t.Dna{}

	d.Add("hello","kemcho")
	n := t.Nucleus{}
	n.Insert(&d)

	c := t.CreateCell()
	c.Insert(&n)

	c.Start()

	fmt.Println("genesis", c)
}
