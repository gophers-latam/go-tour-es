// +build OMIT

package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Buenos días!")
	case t.Hour() < 17:
		fmt.Println("Buenas tardes.")
	default:
		fmt.Println("Buenas noches.")
	}
}
