// +build OMIT

package main

import "fmt"

func main() {
	i, j := 42, 2701

	p := &i         // apunta a la dirección de memoria de i
	fmt.Println(*p) // lee i a través del puntero p
	*p = 21         // establece i a través del puntero p
	fmt.Println(i)  // ve el nuevo valor de i

	p = &j         // apunta a la dirección de memoria de j
	*p = *p / 37   // divide j a través del puntero p
	fmt.Println(j) // ve el nuevo valor de j
}
