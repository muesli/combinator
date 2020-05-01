package main

import (
	"fmt"

	"github.com/muesli/combinator"
)

func main() {
	type User struct {
		Name  string
		Age   uint
		Admin bool
	}

	// make sure the struct's fields share the name and type
	// of the structs you want to generate
	type UserTests struct {
		Name  []string
		Age   []uint
		Admin []bool
	}

	// define potential values
	testData := UserTests{
		Name:  []string{"Alice", "Bob"},
		Age:   []uint{23, 42, 99},
		Admin: []bool{false, true},
	}

	// generate all possible combinations
	var users []User
	combinator.Generate(&users, testData)

	for i, u := range users {
		fmt.Printf("Combination %2d | Name: %-5s | Age: %d | Admin: %v\n", i, u.Name, u.Age, u.Admin)
	}
}
