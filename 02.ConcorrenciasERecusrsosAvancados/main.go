package main

import (
	"context"
	"fmt"
	"net/http"
	"sync"
	"time"
)

/*

- Type Parameters [Generics]
*/

type MeuTipo string

func (MeuTipo) Foo() {}

func main() {
	foo("")
	foo(123)

	var mt MeuTipo = ""
	foo(mt)

	start := time.Now()
	const n = 10
	var wg sync.WaitGroup
	wg.Add(n)

	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	for range 10 {
		go func(ctx context.Context) {
			defer wg.Done()

			resp, err := http.Get("https://google.com")

			if err != nil {
				panic(err)
			}

			defer resp.Body.Close()
			fmt.Println("ok")
		}(ctx)

	}

	wg.Wait()
	fmt.Println(time.Since(start))
}

type MyConstraint interface {
	Foo()
}

func foo[T comparable](arg T) {
	fmt.Println(arg)
}
