package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Name string `json:"name"`
	ID   uint64 `json:"id"`
}

type Animal interface {
	Sound() string
}

type Dog struct{}

func (u User) PrintName() {
	fmt.Println(u.Name)
}

func (u *User) UpdateName(newName string) {
	u.Name = newName
}

type Person struct {
	Name string
}

func (p Person) Greet() {
	fmt.Println("Hello, my name is", p.Name)
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

	dog := Dog{}
	whatDoesThisAnimalSay(dog)

	p := Person{Name: "Alice"}
	p.Greet()
}

func (Dog) Sound() string {
	return "Au! Au!"
}

func whatDoesThisAnimalSay(a Animal) {
	fmt.Println(a.Sound())
}
