package main

import (
	"errors"
	"fmt"
)

func main() {
	a := 10
	b := 0

	res, err := dividir(a, b)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(a, "/", b, res)
}

func dividir(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("nao pode dividir por zero")
	}

	return a / b, nil
}
