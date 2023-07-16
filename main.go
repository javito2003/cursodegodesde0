package main

import (
	"cursodegodesde0/variables"
	"fmt"
)

func main() {
	// variables.RestoVariables()
	estado, texto := variables.ConviertoATexto(1588)
	fmt.Println(estado)
	fmt.Println(texto)
}