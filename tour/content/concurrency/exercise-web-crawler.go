// +build OMIT

package main

import (
	"fmt"
)

type Fetcher interface {
    // Fetch devuelve el cuerpo de la URL y
    // una porción de URL encontradas en esa página.
	Fetch(url string) (body string, urls []string, err error)
}

// Crawl utiliza el buscador para rastrear de forma recursiva
// páginas que comienzan con url, hasta un máximo de profundidad.
func Crawl(url string, depth int, fetcher Fetcher) {
	// TODO: Recuperar URL en paralelo.
    // TODO: No busques la misma URL dos veces.
    // Esta implementación tampoco hace lo siguiente:
	if depth <= 0 {
		return
	}
	body, urls, err := fetcher.Fetch(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("encontrado: %s %q\n", url, body)
	for _, u := range urls {
		Crawl(u, depth-1, fetcher)
	}
	return
}

func main() {
	Crawl("https://golang.org/", 4, fetcher)
}

// fakeFetcher es un Fetcher que devuelve resultados predefinidos.
type fakeFetcher map[string]*fakeResult

type fakeResult struct {
	body string
	urls []string
}

func (f fakeFetcher) Fetch(url string) (string, []string, error) {
	if res, ok := f[url]; ok {
		return res.body, res.urls, nil
	}
	return "", nil, fmt.Errorf("no encontrado: %s", url)
}

// fetcher es un fakeFetcher poblado.
var fetcher = fakeFetcher{
	"https://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"https://golang.org/pkg/",
			"https://golang.org/cmd/",
		},
	},
	"https://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"https://golang.org/",
			"https://golang.org/cmd/",
			"https://golang.org/pkg/fmt/",
			"https://golang.org/pkg/os/",
		},
	},
	"https://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
	"https://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
}
