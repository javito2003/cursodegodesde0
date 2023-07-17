package middlewares

import "fmt"

func sumar(a, b int) int {
	return a + b
}

func restar(a, b int) int {
	return a - b
}

func multiplicar(a, b int) int {
	return a * b
}

func MiMiddleware() {
	fmt.Println("Inicio")

	result := operacionesMidd(sumar)(2, 3)
	fmt.Println(result)

	result2 := operacionesMidd(restar)(10, 6)
	fmt.Println(result2)

	result3 := operacionesMidd(multiplicar)(2, 3)
	fmt.Println(result3)
}

func operacionesMidd(f func(int, int) int) func(int, int) int {
	return func(a, b int) int {
		fmt.Println("Inicio de operacion")
		return f(a, b)
	}
}
