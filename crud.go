package main

//func main() {
//
//	// CREACION DE UNA NUEVA RECETA
//	//r := Receta{
//	//	nombre: "Pie de Durazno",
//	//	tipoPlato: "Postre",
//	//	preparacion: "Mezclar todos los ingredientes y hornear",
//	//	porciones: 15,
//	//}
//	//
//	//_, err := createReceta(r)
//	//if err != nil {
//	//	log.Fatal(err)
//	//}
//
//
//	//ACTUALIZACION DE UNA RECETA
//	//r := Receta{
//	//	idRecetas: 412471344852271105,
//	//	nombre: "Mousse de Mora",
//	//	tipoPlato: "Postre",
//	//	preparacion: "Mezclar todos los ingredientes y refrigerar por 4 horas",
//	//	porciones: 10,
//	//}
//	//
//	//_, err := updateReceta(r)
//	//if err != nil {
//	//	log.Fatal(err)
//	//}
//	//fmt.Println("Receta actualizada")
//
//
//	//ELIMINACION DE UNA RECETA
//	//_, err := deleteReceta(412467458728361985)
//	//if err != nil {
//	//	log.Fatal(err)
//	//}
//	//fmt.Println("Receta eliminada")
//
//
//	// LISTA LAS RECETAS (CON FILTRO POR NOMBRE)
//	listaRecetas, err := allRecetas("")
//	if err != nil {
//		log.Fatal(err)
//	}
//	fmt.Println("\n+++++ RECETAS +++++\n")
//	for i := 0; i < len(listaRecetas); i++ {
//		fmt.Println(listaRecetas[i])
//	}
//
//
//	// LISTA UNA RECETA (POR ID)
//	//recetaSeleccionada, err := getReceta(412783567569125377)
//	//if err != nil {
//	//	log.Fatal(err)
//	//}
//	//fmt.Println("\n+++++ RECETA SELECCIONADA +++++\n")
//	//fmt.Println(recetaSeleccionada)
//
//}