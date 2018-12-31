package main

import (
	"database/sql"
)

// Estructura de datos para Receta
type Receta struct {
	IdRecetas int `json:"idrecetas,omitempty"`
	Nombre string `json:"nombre"`
	TipoPlato string `json:"tipoplato"`
	Preparacion string `json:"preparacion"`
	Porciones int `json:"porciones"`
	Ingredientes string `json:"ingredientes"`
}


// ---------- CRUD PARA RECETAS ----------

// Funcion getReceta: Obtiene los datos de una Receta por medio de su id
// Receiver:
// 		- r *Receta: Elemento tipo Receta, ingresa con el id y sale con todos los campos de la receta
// Input:
// 		- db *sql.DB: Elemento de base de datos de la aplicacion
// Output:
// 		- error: Error generado, si aplica
func (r *Receta) getReceta(db *sql.DB) error {
	return db.QueryRow("SELECT idrecetas, nombre, tipoplato, preparacion, porciones, ingredientes FROM recetas WHERE idrecetas = $1",
		r.IdRecetas).Scan(&r.IdRecetas, &r.Nombre, &r.TipoPlato, &r.Preparacion, &r.Porciones, &r.Ingredientes)
}

// Funcion getRecetas: Obtiene los datos de las Recetas filtradas con base en el nombre, si el texto de busqueda es vacío regresa todas las recetas
// Input:
// 		- db *sql.DB: Elemento de base de datos de la aplicacion
// 		- start int: Posicion inicial del elemento a mostrar en la consulta, para paginacion
// 		- count int: Cantidad de elementos a mostrar por consulta, para paginacion
// 		- txtBusqueda string: Texto para buscar por nombre la receta
// Output:
// 		- recetas []Receta: Lista de elementos tipo Receta
// 		- error: Error generado, si aplica
func getRecetas(db *sql.DB, start, count int, txtBusqueda string) ([]Receta, error) {

	txtBusqueda="%"+txtBusqueda+"%"
	rows, err := db.Query("SELECT * FROM recetas WHERE LOWER(nombre) like LOWER($1) ORDER BY nombre LIMIT $2 OFFSET $3",
		txtBusqueda, count, start)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	recetas := []Receta{}
	for rows.Next() {
		var r Receta
		if err := rows.Scan(&r.IdRecetas, &r.Nombre, &r.TipoPlato, &r.Preparacion, &r.Porciones, &r.Ingredientes); err != nil {
			return nil, err
		}
		recetas = append(recetas, r)
	}

	return recetas, nil
}

// Funcion createReceta: Crea una nueva receta
// Receiver:
// 		- r *Receta: Elemento tipo Receta, ingresa con los campos de la receta a crear sin el id y sale con los campos más el id
// Input:
// 		- db *sql.DB: Elemento de base de datos de la aplicacion
// Output:
// 		- error: Error generado, si aplica
func (r *Receta) createReceta(db *sql.DB) error {
	err := db.QueryRow("INSERT INTO recetas (nombre, tipoplato, preparacion, porciones, ingredientes)	VALUES ($1, $2, $3, $4, $5) RETURNING idrecetas",
		r.Nombre, r.TipoPlato, r.Preparacion, r.Porciones, r.Ingredientes).Scan(&r.IdRecetas)

	if err != nil {
		return err
	}

	return nil
}

// Funcion updateReceta: Actualiza los datos de una receta
// Receiver:
// 		- r *Receta: Elemento tipo Receta, ingresa y sale con los campos a actualizar de la receta
// Input:
// 		- db *sql.DB: Elemento de base de datos de la aplicacion
// Output:
// 		- error: Error generado, si aplica
func (r *Receta) updateReceta(db *sql.DB) error {
	_, err := db.Exec("UPDATE recetas set nombre=$1, tipoplato=$2, preparacion=$3, porciones=$4, ingredientes=$5 where idrecetas=$6",
			r.Nombre, r.TipoPlato, r.Preparacion, r.Porciones, r.Ingredientes ,r.IdRecetas)

	return err
}

// Funcion deleteReceta: Elimina una receta
// Receiver:
// 		- r *Receta: Elemento tipo Receta, ingresa y sale con el id
// Input:
// 		- db *sql.DB: Elemento de base de datos de la aplicacion
// Output:
// 		- error: Error generado, si aplica
func (r *Receta) deleteReceta(db *sql.DB) error {
	_, err := db.Exec("DELETE FROM recetas WHERE idrecetas=$1", r.IdRecetas)

	return err
}



