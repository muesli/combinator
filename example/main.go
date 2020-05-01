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

	c := combinator.Generate(User{}, testData)
	for i, v := range c {
		u := v.(User)
		fmt.Printf("Combination %2d - Name: %-6s - Age: %d - Admin: %v\n", i, u.Name, u.Age, u.Admin)
	}
}
