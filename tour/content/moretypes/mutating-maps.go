//go:build OMIT
// +build OMIT

package main

import "fmt"

func main() {
	m := make(map[string]int)

	m["Respuesta"] = 42
	fmt.Println("El valor:", m["Respuesta"])

	m["Respuesta"] = 48
	fmt.Println("El valor:", m["Respuesta"])

	delete(m, "Respuesta")
	fmt.Println("El valor:", m["Respuesta"])

	v, ok := m["Respuesta"]
	fmt.Println("El valor:", v, "¿Presente?", ok)
}
