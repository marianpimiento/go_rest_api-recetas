package main

import (
	"fmt"
	"log"
)

func main() {
	//r := Receta{
	//	nombre: "Torta de chocolate",
	//	tipoPlato: "Postre",
	//	preparacion: "Mezclar todos los ingredientes y hornear",
	//	porciones: 12,
	//}

	//err := RecetaCrear(r)
	//if err != nil {
	//	log.Fatal(err)
	//}

	var listaRecetas []Receta
	listaRecetas, err2 := RecetaListar("tor")
	if err2 != nil {
		log.Fatal(err2)
	}

	fmt.Println("\n\nCreado exitosamente")
	fmt.Println(listaRecetas)
}