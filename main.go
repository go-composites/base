package main

import (
	"fmt"

	Base "github.com/golang-cop/base/src"
)

func main() {
	o := Base.New()
	fmt.Printf("o.Kind(): %s\n", o.Kind())
	fmt.Printf("o.RespondTo(`Kind`): %t\n", o.RespondTo("Kind"))
	fmt.Printf("o.RespondTo(`Nope`): %t\n", o.RespondTo("Nope"))
	fmt.Printf("o.Methods(): %+v\n", o.Methods())
}
