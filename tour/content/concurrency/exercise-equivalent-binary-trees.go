// +build no-build OMIT

package main

import "golang.org/x/tour/tree"

// Walk recorre el árbol t enviando todos los valores
// del árbol al canal ch.
func Walk(t *tree.Tree, ch chan int)

// Same determina si los árboles
// t1 y t2 contienen los mismos valores.
func Same(t1, t2 *tree.Tree) bool

func main() {
}
