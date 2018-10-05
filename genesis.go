package main

import (
	b "cellula/blueprint"
	i "cellula/info"
	t "cellula/transport"
)

func main() {
	d := t.Dna{}

	d.Add(i.Hello{}, i.Kemcho{}, b.Blueprint{})
	n := t.Nucleus{}
	n.Insert(&d)

	c := t.CreateCell()
	c.Insert(&n)
	c.Start()

}
