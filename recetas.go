package main

import (
	"fmt"
)

// Estructura de la tabla RECETAS
type Receta struct {
	idRecetas int `json:"idrecetas,omitempty"`
	nombre string `json:"nombre"`
	tipoPlato string `json:"tipoplato"`
	preparacion string `json:"preparacion"`
	porciones int `json:"porciones"`
}


// Funcion getReceta: Obtiene los datos de una Receta por medio de su id
// Input:
// 		- idBusqueda int: Id de la receta a buscar
// Output:
// 		- res Receta: Elemento tipo receta con los datos, si no hubo error
// 		- err error: Error generado, si aplica
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

	var idRecetas, porciones int
	var nombre, tipoPlato, preparacion string

	err = db.QueryRow(query, idBusqueda).Scan(&idRecetas, &nombre, &tipoPlato, &preparacion, &porciones)
	if err == nil {
		res = Receta{
			idRecetas: idRecetas,
			nombre: nombre,
			tipoPlato: tipoPlato,
			preparacion: preparacion,
			porciones: porciones,
		}
	}

	return res, err
}


// Funcion allRecetas: Obtiene los datos de las Recetas filtradas con base en el nombre, si el texto de busqueda es vacío regresa todas las recetas
// Input:
// 		- txtBusqueda string: Texto para buscar por nombre la receta
// Output:
// 		- recetas []Receta: Lista de elementos tipo receta con los datos, si no hubo error
// 		- err error: Error generado, si aplica
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
			var idRecetas, porciones int
			var nombre, tipoPlato, preparacion string

			err = rows.Scan(&idRecetas, &nombre, &tipoPlato, &preparacion, &porciones)
			if err == nil {
				recetaActual := Receta{
					idRecetas: idRecetas,
					nombre: nombre,
					tipoPlato: tipoPlato,
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


// Funcion insertReceta: Crea una nueva receta
// Input:
// 		- r Receta: Elemento tipo Receta con los datos para la creacion
// Output:
// 		- recetaID int: Id de la nueva receta creada si no hubo error, de lo contrario es cero
// 		- err error: Error generado, si aplica
func insertReceta(r Receta) (int, error) {

	var recetaID int

	query := `INSERT INTO recetas (nombre, tipoplato, preparacion, porciones)	VALUES ($1, $2, $3, $4) RETURNING idrecetas`
	db := getConnection()
	defer db.Close()

	stmt, err := db.Prepare(query)
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	err = db.QueryRow(query, r.nombre, r.tipoPlato , r.preparacion , r.porciones).Scan(&recetaID)
	if err != nil {
		return 0, err
	}
	fmt.Println("ID receta:", recetaID)

	return recetaID, nil
}


// Funcion updateReceta: Actualizar una receta
// Input:
// 		- r Receta: Elemento tipo Receta con los datos para la actualizacion
// Output:
// 		- rowsUpdated int: Cantidad de filas actualizadas si no hubo error, de lo contrario es cero
// 		- err error: Error generado, si aplica
func updateReceta(r Receta) (int, error) {

	query := `UPDATE recetas set nombre=$1, tipoplato=$2, preparacion=$3, porciones=$4 where idrecetas=$5 RETURNING idrecetas`

	db := getConnection()
	defer db.Close()

	stmt, err := db.Prepare(query)
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	res, err := db.Exec(query, r.nombre, r.tipoPlato , r.preparacion , r.porciones, r.idRecetas)
	if err != nil {
		return 0, err
	}

	rowsUpdated, err := res.RowsAffected()
	if err != nil {
		return 0, err
	}

	return int(rowsUpdated), err
}


// Funcion deleteReceta: Eliminar una receta
// Input:
// 		- recetaID int: Id de la receta a eliminar
// Output:
// 		- rowsDeleted int:  Cantidad de filas eliminadas si no hubo error, de lo contrario es cero
// 		- err error: Error generado, si aplica
func deleteReceta(recetaID int) (int, error) {

	query := `DELETE FROM recetas WHERE idrecetas=$1`

	db := getConnection()
	defer db.Close()

	stmt, err := db.Prepare(query)
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	res, err := db.Exec(query, recetaID)
	if err != nil {
		return 0, err
	}

	rowsDeleted, err := res.RowsAffected()
	if err != nil {
		return 0, err
	}

	return int(rowsDeleted), nil

}