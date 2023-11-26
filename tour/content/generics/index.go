//go:build OMIT
// +build OMIT

package main

import "fmt"

// Index devuelve el índice de x en s, o -1 si no se encuentra.
func Index[T comparable](s []T, x T) int {
	for i, v := range s {
        // v y x son del tipo T, que tiene la restricción
        // comparable, por lo que podemos usar == aquí.
		if v == x {
			return i
		}
	}
	return -1
}

func main() {
    // Index funciona en un slice de ints
    si := []int{10, 20, 15, -10}
	fmt.Println(Index(si, 15))

    // Index también funciona en un slice de enteros
    ss := []string{"foo", "bar", "baz"}
	fmt.Println(Index(ss, "hola"))
}
