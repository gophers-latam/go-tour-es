//go:build OMIT
// +build OMIT

package main

import "fmt"

const Pi = 3.14

func main() {
	const World = "世界"
	fmt.Println("Hola", World)
	fmt.Println("Feliz día de", Pi)

	const Truth = true
	fmt.Println("Go manda?", Truth)
}
