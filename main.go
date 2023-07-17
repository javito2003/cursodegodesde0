package main

import (
	"cursodegodesde0/goroutines"
)

func main() {
	// numero, mensaje := ejercicios.EsMayorA100("50")
	// fmt.Println(numero)
	// fmt.Println(mensaje)

	// teclado.IngresoNumeros()
	// iteraciones.Iterar()
	// ejercicios.IngresaNumeroParaMultiplicar()
	// files.SumaTabla()
	// files.LeoArchivo()
	// funciones.Calculos()
	// funciones.LlamarClosure()
	// funciones.Exponencia(2)
	// arreglos_slices.MuestroSlices()
	// arreglos_slices.Capacidad()
	// mapas.MostrarMapa()
	// users.AltaUsuario()

	// Pedro := new(modelos.Hombre)
	// ejer_interfaces.HumanosRespirando(Pedro)

	// Martina := new(modelos.Mujer)
	// ejer_interfaces.HumanosRespirando(Martina)

	// defer_ejemplo.EjemploPanic()

	canal1 := make(chan bool)
	go goroutines.MiNombreLento("Javico", canal1)
	<-canal1

	// fmt.Println("Estoy aqui")
	// var x string
	// fmt.Scanln(&x)
}
