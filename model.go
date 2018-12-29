package main

import (
	"database/sql"
)

// STRUCT DEFINIDO EN RECETAS.GO
//type Receta struct {
//	idRecetas int `json:"idrecetas,omitempty"`
//	nombre string `json:"nombre"`
//	tipoPlato string `json:"tipoplato"`
//	preparacion string `json:"preparacion"`
//	porciones int `json:"porciones"`
//}


func (r *Receta) getRecetaModel(db *sql.DB) error {
	return db.QueryRow("SELECT * FROM recetas WHERE idrecetas = $1",
		r.idRecetas).Scan(&r.nombre, &r.tipoPlato, &r.preparacion, &r.porciones)
}

func (r *Receta) updateRecetaModel(db *sql.DB) error {
	_, err :=
		db.Exec("UPDATE recetas set nombre=$1, tipoplato=$2, preparacion=$3, porciones=$4 where idrecetas=$5",
			r.nombre, r.tipoPlato, r.preparacion, r.porciones, r.idRecetas)

	return err
}

func (r *Receta) deleteRecetaModel(db *sql.DB) error {
	_, err := db.Exec("DELETE FROM recetas WHERE idrecetas=$1", r.idRecetas)

	return err
}

func (r *Receta) createRecetaModel(db *sql.DB) error {
	err := db.QueryRow(
		"INSERT INTO recetas (nombre, tipoplato, preparacion, porciones)	VALUES ($1, $2, $3, $4) RETURNING idrecetas",
		r.nombre, r.tipoPlato, r.preparacion, r.porciones).Scan(&r.idRecetas)

	if err != nil {
		return err
	}

	return nil
}

func getRecetasModel(db *sql.DB, start, count int, txtBusqueda string) ([]Receta, error) {
	rows, err := db.Query(
		"SELECT * FROM recetas WHERE LOWER(nombre) like LOWER($1) ORDER BY nombre LIMIT $2 OFFSET $3",
		txtBusqueda, count, start)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	recetas := []Receta{}

	for rows.Next() {
		var r Receta
		if err := rows.Scan(&r.idRecetas, &r.nombre, &r.tipoPlato, &r.preparacion, &r.porciones); err != nil {
			return nil, err
		}
		recetas = append(recetas, r)
	}

	return recetas, nil
}