package main

import (
	"fmt"
	"log"
	"reflect"
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
	listaRecetas, err := ListarRecetas("")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("\n+++++ RECETAS +++++\n")
	for i := 0; i < len(listaRecetas); i++ {
		fmt.Println(listaRecetas[i])
	}



	//var recetaSeleccionada Receta
	recetaSeleccionada, err := getReceta2(1)
	fmt.Println(reflect.TypeOf(err))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("\n+++++ RECETA SELECCIONADA +++++\n")

	fmt.Println(recetaSeleccionada)


}