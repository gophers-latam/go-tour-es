// +build no-build OMIT

package main

import "golang.org/x/tour/reader"

type MyReader struct{}

// TODO: Agregar un metodo Read([]byte) (int, error) a MyReader.

func main() {
	reader.Validate(MyReader{})
}
