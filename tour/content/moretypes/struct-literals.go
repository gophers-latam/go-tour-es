// +build OMIT

package main

import "fmt"

type Vertex struct {
	X, Y int
}

var (
	v1 = Vertex{1, 2}  // tiene tipo Vertex
	v2 = Vertex{X: 1}  // Y:0 es implicito
	v3 = Vertex{}      // X:0 y Y:0
	p  = &Vertex{1, 2} // tiene tipo *Vertex
)

func main() {
	fmt.Println(v1, p, v2, v3)
}
