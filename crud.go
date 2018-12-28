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
	listaRecetas, err2 := RecetaListar("")
	if err2 != nil {
		log.Fatal(err2)
	}

	fmt.Println("\n+++++ RECETAS +++++\n")


	for i := 0; i < len(listaRecetas); i++ {
		fmt.Println(listaRecetas[i])
	}
}