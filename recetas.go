package main

import (
	"fmt"
)

// Estructura de la tabla RECETAS
type Receta struct {
	idRecetas int `json:"idrecetas,omitempty"`
	nombre string `json:"nombre,omitempty"`
	tipoPlato string `json:"tipoplato,omitempty"`
	preparacion string `json:"preparacion,omitempty"`
	porciones int `json:"porciones,omitempty"`
}

// CRUD

// Crear RECETAS
func RecetaCrear(r Receta) error {
	query := `INSERT INTO recetas (nombre, tipoplato, preparacion, porciones)
	VALUES ($1, $2, $3, $4) RETURNING idrecetas`
	db := getConnection()
	defer db.Close()

	stmt, err := db.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	var idreceta int

	err = db.QueryRow(query, r.nombre, r.tipoPlato , r.preparacion , r.porciones).Scan(&idreceta)
	if err != nil {
		return err
	}
	fmt.Println("ID receta:", idreceta)

	return nil
}


func RecetaListar(txtBusqueda string) ([]Receta, error) {

	var recetas []Receta

	query := `SELECT * FROM recetas WHERE LOWER(nombre) like LOWER($1)`

	db := getConnection()
	defer db.Close()

	stmt, err := db.Prepare(query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	txtBusqueda="%"+txtBusqueda+"%"
	rowsRecetas, err := db.Query(query, txtBusqueda)

	if err != nil {
		return nil, err
	}
	defer rowsRecetas.Close()

	//fmt.Println("\n+++++ RECETAS +++++\n")

	for rowsRecetas.Next() {
		var id, porciones int
		var nombre, tipo, preparacion string
		if err := rowsRecetas.Scan(&id, &nombre, &tipo, &preparacion, &porciones); err != nil {
			return nil, err
		}
		//fmt.Printf("Receta: %s\n", nombre)
		//fmt.Printf("Tipo: %s\n", tipo)
		//fmt.Printf("Porciones: %d\n", porciones)
		//fmt.Printf("Preparacion:\n %s\n\n", preparacion)

		recetaActual := Receta{
			idRecetas: id,
			nombre: nombre,
			tipoPlato: tipo,
			preparacion: preparacion,
			porciones: porciones,
		}

		recetas = append(recetas, recetaActual)
		//fmt.Println(recetaActual)
	}
	//fmt.Println(recetas)
	return recetas, nil

}