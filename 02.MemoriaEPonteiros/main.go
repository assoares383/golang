package main

import "fmt"

func main() {
	x := 10
	// Ponteiro para x
	// p := &x

	take(&x)

	fmt.Println(x)

	arr := [5]int{1, 2, 3, 4, 5}
	slice := arr[1:4]
	arr[2] = 15
	slice[0] = 123
	fmt.Println(slice)
}

func take(x *int) {
	*x = 100
}
