package main

import "fmt"

type Foo struct {
	Bar string
	Baz float64
	Quz []string
}

func main() {
	f := Foo{
		Bar: "",
		Baz: 0,
		Quz: []string{},
	}

	fmt.Println(f)
}
