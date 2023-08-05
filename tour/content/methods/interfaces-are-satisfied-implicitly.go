// +build OMIT

package main

import "fmt"

type I interface {
	M()
}

type T struct {
	S string
}

// Este método significa que el tipo T implementa la interface I,
// Pero no necesitamos declarar explícitamente lo que hace.
func (t T) M() {
	fmt.Println(t.S)
}

func main() {
	var i I = T{"Hola"}
	i.M()
}
