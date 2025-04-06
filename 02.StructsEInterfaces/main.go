package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Name string `json:"name"`
	ID   uint64 `json:"id"`
}

func (u User) PrintName() {
	fmt.Println(u.Name)
}

func (u *User) UpdateName(newName string) {
	u.Name = newName
}

func main() {
	// Uma estrutura e uma colecao de campos
	user := User{Name: "Alexandre", ID: 10}

	// PrintName
	fmt.Println(user.Name)
	user.PrintName()

	// UpdateName
	user.UpdateName("Lethicia")
	fmt.Println(user)

	res, err := json.Marshal(user)

	if err != nil {
		panic(err)
	}

	fmt.Println(string(res))
}
