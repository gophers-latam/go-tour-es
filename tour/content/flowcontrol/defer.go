// +build OMIT

package main

import "fmt"

func main() {
	defer fmt.Println("mundo")

	fmt.Println("hola")
}
