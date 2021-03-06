package main

import "fmt"

type N int

func (n N) value() {
	n++
	fmt.Printf("%p %v\n", &n, n)
}

func (n *N) pointer() {
	*n++
	fmt.Printf("%p %v\n", n, *n)
}

func main() {
	var n N = 1
	fmt.Printf("%p %v\n", &n, n)

	f1 := N.value
	f2 := (*N).value
	//f3 := N.pointer // N.pointer undefined (type N has no method pointer)
	f4 := (*N).pointer

	f1(n)
	fmt.Printf("%p %v\n", &n, n)
	f2(&n)
	fmt.Printf("%p %v\n", &n, n)
	// f3(&n)
	// fmt.Printf("%p %v\n", &n, n)
	f4(&n)
	fmt.Printf("%p %v\n", &n, n)

	N.value(n)
	fmt.Printf("%p %v\n", &n, n)
	// (*N).pointer(n)
    // cannot use n (type N) as type *N in argument to (*N).pointer
}
