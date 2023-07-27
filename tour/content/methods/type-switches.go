// +build OMIT

package main

import "fmt"

func do(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("Dos veces %v es %v\n", v, v*2)
	case string:
		fmt.Printf("%q tiene %v bytes de largo\n", v, len(v))
	default:
		fmt.Printf("No entiendo nada sobre tipos %T!\n", v)
	}
}

func main() {
	do(21)
	do("hello")
	do(true)
}
