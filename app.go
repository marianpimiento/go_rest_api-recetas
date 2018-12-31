package main

import (
	"database/sql"
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"

	_ "github.com/lib/pq"
)


type App struct {
	Router *mux.Router
	DB     *sql.DB
}

func (a *App) Initialize() {
	dsn := "postgresql://maxroach@localhost:26257/recetas?sslmode=disable"

	var err error
	a.DB, err = sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal(err)
	}
	a.Router = mux.NewRouter()
	a.initializeRoutes()
}

func (a *App) Run(addr string) {
	log.Fatal(http.ListenAndServe(addr, a.Router))

	defer a.DB.Close()
}


func (a *App) initializeRoutes() {
	a.Router.HandleFunc("/recetas", a.getRecetas).Methods("GET")
	a.Router.HandleFunc("/receta", a.createReceta).Methods("POST")
	a.Router.HandleFunc("/receta/{id:[0-9]+}", a.getReceta).Methods("GET")
	a.Router.HandleFunc("/receta/{id:[0-9]+}", a.updateReceta).Methods("PUT")
	a.Router.HandleFunc("/receta/{id:[0-9]+}", a.deleteReceta).Methods("DELETE")
}


//-----------------------
func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}



//--------------------- Handlers
func (a *App) getReceta(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Id de receta invalido")
		return
	}

	rec := Receta{IdRecetas: id}
	if err := rec.getRecetaModel(a.DB); err != nil {

		switch err {
		case sql.ErrNoRows:
			respondWithError(w, http.StatusNotFound, "Receta no encontrada")
		default:
			respondWithError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}

	respondWithJSON(w, http.StatusOK, rec)
}

func (a *App) getRecetas(w http.ResponseWriter, r *http.Request) {
	count, _ := strconv.Atoi(r.FormValue("count"))
	start, _ := strconv.Atoi(r.FormValue("start"))
	txtBusqueda := r.FormValue("txtBusqueda")

	if count > 10 || count < 1 {
		count = 10
	}
	if start < 0 {
		start = 0
	}

	recetas, err := getRecetasModel(a.DB, start, count, txtBusqueda)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, recetas)
}

func (a *App) createReceta(w http.ResponseWriter, r *http.Request) {
	var rec Receta
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&rec); err != nil {
		respondWithError(w, http.StatusBadRequest, "Request de la carga util no valido")
		return
	}
	defer r.Body.Close()

	if err := rec.createRecetaModel(a.DB); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusCreated, rec)
}


func (a *App) updateReceta(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Id de receta invalido")
		return
	}

	var rec Receta
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&rec); err != nil {
		respondWithError(w, http.StatusBadRequest, "Request de la carga util no valido")
		return
	}
	defer r.Body.Close()
	rec.IdRecetas = id

	if err := rec.updateRecetaModel(a.DB); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, rec)
}

func (a *App) deleteReceta(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Id de receta invalido")
		return
	}

	rec := Receta{IdRecetas: id}
	if err := rec.deleteRecetaModel(a.DB); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, map[string]string{"result": "success"})
}