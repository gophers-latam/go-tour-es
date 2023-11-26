package main

// List representa una lista enlazada individualmente que contiene
// valores de cualquier tipo.
type List[T any] struct {
	next *List[T]
	val  T
}

func main() {
}
