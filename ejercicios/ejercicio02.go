package ejercicios

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func IngresaNumeroParaMultiplicar() string {
	scanner := bufio.NewScanner(os.Stdin)
	var numero int
	var err error
	var texto string

	fmt.Println("Ingrese un numero:")
	for {
		if scanner.Scan() {
			numero, err = strconv.Atoi(scanner.Text())
			if err == nil {
				break
			}
			fmt.Println("Ingresa un valor valido")
		}
	}

	for i := 1; i <= 10; i++ {
		resultado := numero * i
		texto += fmt.Sprintf("%d x %d = %d \n", numero, i, resultado)
	}

	return texto
}
