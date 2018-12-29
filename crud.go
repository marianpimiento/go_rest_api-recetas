package main

import (
	"fmt"
	"log"
	"reflect"
)

func main() {
	//r := Receta{
	//	nombre: "Pie de Manzana",
	//	tipoPlato: "Postre",
	//	preparacion: "Mezclar todos los ingredientes y hornear",
	//	porciones: 15,
	//}
	//
	//err := RecetaCrear(r)
	//if err != nil {
	//	log.Fatal(err)
	//}


	//var listaRecetas []Receta
	listaRecetas, err := allRecetas("")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("\n+++++ RECETAS +++++\n")
	for i := 0; i < len(listaRecetas); i++ {
		fmt.Println(listaRecetas[i])
	}


	//var recetaSeleccionada Receta
	recetaSeleccionada, err := getReceta(412783567569125377)
	fmt.Println(reflect.TypeOf(err))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("\n+++++ RECETA SELECCIONADA +++++\n")

	fmt.Println(recetaSeleccionada)


}