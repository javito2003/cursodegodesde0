package files

import (
	"cursodegodesde0/ejercicios"
	"fmt"
	"io/ioutil"
	"os"
)

var filename string = "./files/archivos/tabla.txt"

func GrabaTabla() {
	var texto string = ejercicios.IngresaNumeroParaMultiplicar()
	archivo, err := os.Create(filename)
	if err != nil {
		panic("Error al crear archivo, " + err.Error())
	}
	fmt.Fprintln(archivo, texto)
	archivo.Close()
}

func SumaTabla() {
	var texto string = ejercicios.IngresaNumeroParaMultiplicar()
	if !Append(texto) {
		fmt.Println("Error al concatenar contenido")
	}
}

func Append(texto string) bool {
	archivo, err := os.OpenFile(filename, os.O_WRONLY|os.O_APPEND, 0644) 
	if err != nil {
		fmt.Println("Error durante el Append " + err.Error())
		return false
	}
	_, err = archivo.WriteString(texto)
	if err != nil {
		fmt.Println("Error durante el Append " + err.Error())
	}
	archivo.Close()

	return true
}

func LeoArchivo() {
	archivo, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Error al leer archivo " + err.Error())
		return
	}
	fmt.Println(string(archivo))
}
