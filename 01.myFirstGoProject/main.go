package main

import (
	"fmt"
	"math"
	"time"
)

// bool
//
// int int8 int16 int32 int64
// uint uint8 uint16 uint32 uint64 uintptr
//
// byte = uint8
//
// rune = int32
//
// float32 float64
//
// complex64 complex128
//
// string

func main() {
	// This is my first Go project
	fmt.Println("Hello, World!")

	var x bool = true
	fmt.Println(x)

	var b uint8 = 10
	takeBite(b)

	// Array
	arr := [5]int{1, 2, 3, 4, 5}
	fmt.Println(arr)

	// Loop
	var res int
	for i := 0; i < 10; i++ {
		res++
	}
	fmt.Println(res)

	// Loop com range
	for i := range 10 {
		fmt.Println(i)
	}

	// If
	if x := math.Sqrt(4); x < 1 {
		fmt.Println(x)
	} else if x < 1 {
		fmt.Println("maior que zero")
	} else {
		fmt.Println("caiu no else")
	}

	// Switches
	do(1)
	do(2)
	do(3)
}

func takeBite(b byte) {
	fmt.Println(b)
}

func do(x int) {
	switch x {
	case 1:
		fmt.Println(1)
	case 2:
		fmt.Println(2)
	default:
		fmt.Println("outra coisa")
	}

	fmt.Println(isWeekend(time.Now()))
}

func isWeekend(x time.Time) bool {
	switch {
	case x.Weekday() > 0 && x.Weekday() < 6:
		return false
	default:
		return true
	}
}
