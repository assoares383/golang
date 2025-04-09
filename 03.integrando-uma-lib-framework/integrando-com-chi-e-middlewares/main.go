package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
)

func main() {
	r := chi.NewMux()

	r.Get("/horario", func(w http.ResponseWriter, r *http.Request) {
		now := time.Now()
		fmt.Fprintln(w, now)
	})

	if err := http.ListenAndServe(":8080", r); err != nil {
		panic(err)
	}
}
