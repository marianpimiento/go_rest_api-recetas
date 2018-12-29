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


func getReceta(idBusqueda int) (Receta, error) {

	res := Receta{}

	query := `SELECT * FROM recetas WHERE idrecetas = $1`

	db := getConnection()
	defer db.Close()

	stmt, err := db.Prepare(query)
	if err != nil {
		return res, err
	}
	defer stmt.Close()

	var id, porciones int
	var nombre, tipo, preparacion string

	err = db.QueryRow(query, idBusqueda).Scan(&id, &nombre, &tipo, &preparacion, &porciones)
	if err == nil {
		res = Receta{
			idRecetas: id,
			nombre: nombre,
			tipoPlato: tipo,
			preparacion: preparacion,
			porciones: porciones,
		}
	}

	return res, err
}


func allRecetas(txtBusqueda string) ([]Receta, error) {

	recetas := []Receta{}

	query := `SELECT * FROM recetas WHERE LOWER(nombre) like LOWER($1) ORDER BY nombre;`

	db := getConnection()
	defer db.Close()

	stmt, err := db.Prepare(query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	txtBusqueda="%"+txtBusqueda+"%"
	rows, err := db.Query(query, txtBusqueda)
	defer rows.Close()


	if err == nil {
		for rows.Next() {
			var id, porciones int
			var nombre, tipo, preparacion string

			err = rows.Scan(&id, &nombre, &tipo, &preparacion, &porciones)
			if err == nil {
				recetaActual := Receta{
					idRecetas: id,
					nombre: nombre,
					tipoPlato: tipo,
					preparacion: preparacion,
					porciones: porciones,
				}

				recetas = append(recetas, recetaActual)

			} else {
				return recetas, err
			}
		}
	} else {
		return recetas, err
	}

	return recetas, err
}
