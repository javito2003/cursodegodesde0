package arreglos_slices

import "fmt"

var tabla [10]int = [10]int { 10, 0, 59 }
var matriz [20][30]int 

func MuestroArreglos() {
	tabla[7] = 33
	tabla[3] = 2

	tabla2 := [10]string { "Raul", "Javier", "Moreno" }
	fmt.Println(tabla)
	fmt.Println(tabla2)

	for i:= 0; i<len(tabla); i++ {
		fmt.Println(tabla[i])
	}

	matriz[7][24] = 50
	fmt.Println(matriz)
}
