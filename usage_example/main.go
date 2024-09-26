package main

import (
	"fmt"
	"plugin"
)

func main() {
	fmt.Println("outer")

	p, err := plugin.Open("../plug/plug.so")
	if err != nil {
		panic(err)
	}

	g, err := p.Lookup("Greet")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%#v\n", g)
	greet := g.(func(string) string)

	fmt.Println(greet("Anne"))
}
