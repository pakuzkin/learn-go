package main

import (
	"fmt"
)

type person struct {
	first string
	last  string
}

func pprint(a any) {
	fmt.Printf("%v\t%T\n", a, a)
}

func incVal(val int) int {
	val++
	fmt.Println("incVal", &val, val) // will have different mem address than passed in parameter because its a copy
	return val
}

func incPtr(p *int) {
	*p++
	fmt.Println("incPtr", *p, p)
}

func renameVal(p person) {
	p.first = "Val"
}

func renamePtr(p *person) {
	p.first = "Ptr"
}

func main() {
	var x int
	x = 42

	pprint(x)
	pprint(&x)

	a := 10
	b := 20

	incVal(a)
	incPtr(&b)

	fmt.Println("a does not change", a, &a)
	fmt.Println("b does change", b, &b)

	h1 := person{
		first: "John",
		last:  "Doe",
	}

	h2 := person{
		first: "Marry",
		last:  "Jane",
	}

	renameVal(h1)
	renamePtr(&h2)

	fmt.Println("h1 does not change", h1, &h1)
	fmt.Println("h2 does change", h2, &h2)

}
