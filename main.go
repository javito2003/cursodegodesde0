package main

import (
	"fmt"
	"runtime"
)

func main() {
	// variables.RestoVariables()
	// estado, texto := variables.ConviertoATexto(1588)
	// fmt.Println(estado)
	// fmt.Println(texto)
	
	os := runtime.GOOS
	// if os := runtime.GOOS; os == "Linux." 
	if os == "linux" {
		fmt.Println("Soy linux")
	} else {
		fmt.Println("Soy Windows")
	}

	switch os {
	case "linux":
		fmt.Println("Esto es linux")
	case "darwin":
		fmt.Println("Esto es darwin")
	default:
		fmt.Printf("%s \n", os)
	}
}