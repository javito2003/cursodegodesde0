package funciones

import "fmt"

func Exponencia(valor int) {
	if valor > 100_000_000 {
		return
	}
	fmt.Println(valor)
	Exponencia(valor * 2)
}
