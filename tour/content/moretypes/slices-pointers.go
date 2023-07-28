// +build OMIT

package main

import "fmt"

func main() {
	nombres := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(nombres)

	a := nombres[0:2]
	b := nombres[1:3]
	fmt.Println(a, b)

	b[0] = "XXX"
	fmt.Println(a, b)
	fmt.Println(nombres)
}
