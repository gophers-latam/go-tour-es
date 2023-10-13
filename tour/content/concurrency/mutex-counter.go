//go:build OMIT
// +build OMIT

package main

import (
	"fmt"
	"sync"
	"time"
)

// SafeCounter es seguro de usar simultáneamente.
type SafeCounter struct {
	mu sync.Mutex
	v  map[string]int
}

// Inc incrementa el contador para la clave dada.
func (c *SafeCounter) Inc(key string) {
	c.mu.Lock()
	// Bloquear para que solo una rutina a la vez pueda acceder al mapa c.v.
	c.v[key]++
	c.mu.Unlock()
}

// Valor devuelve el valor actual del contador para la clave dada.
func (c *SafeCounter) Valor(key string) int {
	c.mu.Lock()
	// Lock so only one goroutine at a time can access the map c.v.
	defer c.mu.Unlock()
	return c.v[key]
}

func main() {
	c := SafeCounter{v: make(map[string]int)}
	for i := 0; i < 1000; i++ {
		go c.Inc("somekey")
	}

	time.Sleep(time.Second)
	fmt.Println(c.Valor("somekey"))
}
