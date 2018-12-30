package main

import (
	"database/sql"
)

// STRUCT DEFINIDO EN RECETAS.GO
type Receta struct {
	IdRecetas int `json:"idrecetas,omitempty"`
	Nombre string `json:"nombre"`
	TipoPlato string `json:"tipoplato"`
	Preparacion string `json:"preparacion"`
	Porciones int `json:"porciones"`
}


func (r *Receta) getRecetaModel(db *sql.DB) error {
	//fmt.Println("getRecetaModel")
	//
	//err := db.QueryRow("SELECT * FROM recetas WHERE idrecetas = $1",
	//	r.IdRecetas).Scan(&r.IdRecetas, &r.Nombre, &r.TipoPlato, &r.Preparacion, &r.Porciones)
	//if err == nil {
	//	fmt.Println("Error")
	//}
	//
	//fmt.Println("Despues metodo largo")
	//
	//fmt.Println(*r)
	//fmt.Println(reflect.TypeOf(r))
	//fmt.Println(reflect.TypeOf(*r))


	return db.QueryRow("SELECT * FROM recetas WHERE idrecetas = $1",
		r.IdRecetas).Scan(&r.IdRecetas, &r.Nombre, &r.TipoPlato, &r.Preparacion, &r.Porciones)
}

func (r *Receta) updateRecetaModel(db *sql.DB) error {
	_, err :=
		db.Exec("UPDATE recetas set nombre=$1, tipoplato=$2, preparacion=$3, porciones=$4 where idrecetas=$5",
			r.Nombre, r.TipoPlato, r.Preparacion, r.Porciones, r.IdRecetas)

	return err
}

func (r *Receta) deleteRecetaModel(db *sql.DB) error {
	_, err := db.Exec("DELETE FROM recetas WHERE idrecetas=$1", r.IdRecetas)

	return err
}

func (r *Receta) createRecetaModel(db *sql.DB) error {
	err := db.QueryRow(
		"INSERT INTO recetas (nombre, tipoplato, preparacion, porciones)	VALUES ($1, $2, $3, $4) RETURNING idrecetas",
		r.Nombre, r.TipoPlato, r.Preparacion, r.Porciones).Scan(&r.IdRecetas)

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
		if err := rows.Scan(&r.IdRecetas, &r.Nombre, &r.TipoPlato, &r.Preparacion, &r.Porciones); err != nil {
			return nil, err
		}
		recetas = append(recetas, r)
	}

	return recetas, nil
}