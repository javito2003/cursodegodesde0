package ejercicios

import "strconv"

func EsMayorA100(strNumero string) (int, string) {
	numero, _ := strconv.Atoi(strNumero)

	var mensaje string
	if (numero > 100) {
		mensaje = "Es mayor a 100"
	} else {
		mensaje = "Es menor a 100" 
	}

	return numero, mensaje
}