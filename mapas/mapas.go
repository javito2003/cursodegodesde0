package mapas

import "fmt"

func MostrarMapa() {
	paises := make(map[string]string)

	paises["Mexico"] = "D.F"
	paises["Argentina"] = "Buenos Aires"
	fmt.Println(paises)
	fmt.Println(paises["Argentina"])

	campeonato := map[string]int {
		"Barcelona": 39,
		"Real Madrid": 38,
		"Chivas": 37,
		"Boca Juniors": 30,
	}

	fmt.Println(campeonato)

	for equipo, puntaje := range campeonato {
		fmt.Println(equipo, puntaje)
	}

	delete(campeonato, "Real Madrid")
	fmt.Println(campeonato)

	puntaje, existe := campeonato["Chivas"]
	fmt.Printf("El puntaje capturado es %d, y el equipo existe = %t", puntaje, existe)
}
