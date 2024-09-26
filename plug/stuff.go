package main

import (
	"fmt"
	"strings"
)

func Greet(name string) string {
	return fmt.Sprintf("Hello %s.", name)
}

func Shout(name string) string {
	return fmt.Sprintf("HI THERE %s!", strings.ToUpper(name))
}
