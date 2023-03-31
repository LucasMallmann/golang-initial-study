package main

import "fmt"

func incr(p *int) int {
	*p++ // increments whats p points to; does not change p
	return *p
}

func f() *int {
	v := 1
	return &v
}

func main() {
	fmt.Println(f() == f())
	v := 1
	incr(&v)
	fmt.Println(incr(&v))

	val := 42
	p := &val
	fmt.Println(*p)
}
