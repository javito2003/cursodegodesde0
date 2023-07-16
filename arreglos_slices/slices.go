package arreglos_slices

import "fmt"

var tablaSlice []int = []int{50, 20, 30}
var arreglo [10]int = [10]int { 1, 2, 3, 4, 5, 6, 7, 8, 9, 10 }

func MuestroSlices() {
	fmt.Println(tablaSlice)

	porcion := arreglo[3:] // Slice creado desde un vector, desde la posicion 3 hasta el ultimo
	porcion2 := arreglo[:5] // Slice creado desde un vector, desde la posicion 0 hasta la 5
	porcion3 := arreglo[6:9]

	fmt.Println(porcion)
	fmt.Println(porcion2)
	fmt.Println(porcion3)

}

func Capacidad() {
	elementos := make([]int, 5, 20)
	fmt.Printf("Largo %d, Capacidad %d \n", len(elementos), cap(elementos))

	nums := make([]int, 0, 0)
	for i := 0; i < 100; i++ {
		nums = append(nums, i)
	}
	fmt.Printf("Largo %d, Capacidad %d", len(nums), cap(nums))
}