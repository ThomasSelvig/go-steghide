// https://gobyexample.com/ is a good reference
package main

import "fmt"

func test() {
	var a, b int = 1, 2
	c := 3  // shorthand initialization var
	b, a = a, b
	fmt.Println(a, b, c)

	// for loops for iterating range
	for i := 1; i < 3; i++ {
	}
	for i := range 3 {
		fmt.Print(i)
	}
	fmt.Println()
	// string loop
	for i, runeVal := range "abcdef" {
		fmt.Printf("%d:%c ", i, runeVal)
	}

	// for loops are also while loops. They can take no condition (while true) or a condition.
	// they accept "continue" and "break" as well.
}
