// Copyright 2011 The Go Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package wc // import "golang.org/x/tour/wc"

import "fmt"

// Test ejecuta un conjunto de pruebas contra f.
func Test(f func(string) map[string]int) {
    ok := true
    for _, c := range testCases {
        got := f(c.in)
        if len(c.want) != len(got) {
            ok = false
        } else {
            for k := range c.want {
                if c.want[k] != got[k] {
                    ok = false
                }
            }
        }
        if !ok {
            fmt.Printf("INCORRECTO\n f(%q) =\n  %#v\n DESEADO:\n  %#v",
            c.in, got, c.want)
            break
        }
        fmt.Printf("CORRECTO\n f(%q) = \n  %#v\n", c.in, got)
    }
}


var testCases = []struct {
    in   string
    want map[string]int
}{
    {"¡Yo estoy aprendiendo Go!", map[string]int{
        "Yo": 1, "estoy": 1, "aprendiendo": 1, "Go!": 1,
    }},
    {"El rápido zorro marrón saltó sobre el perro perezoso.", map[string]int{
        "El": 1, "rápido": 1, "zorro": 1, "marrón": 1, "saltó": 1,
        "sobre": 1, "el": 1, "perro": 1, "perezoso.": 1,
    }},
    {"Me comí una dona. Luego me comí otra dona.", map[string]int{
        "Me": 2, "comí": 2, "una": 1, "dona.": 2, "Luego": 1, "otra": 1,
    }},
    {"Un hombre, un plan, un canal, Panamá.", map[string]int{
        "Un": 1, "hombre": 1, "un": 2, "plan": 1, "canal": 1, "Panamá.": 1,
    }},
}
