package main

import (
	"database/sql"
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"log"
	"net/http"
	"strconv"

	_ "github.com/lib/pq"
)

// Estructura de datos App
type App struct {
	Router *mux.Router
	DB     *sql.DB
}

// Funcion Initialize: Establece conexion con la base de datos e inicializa el router
// Receiver:
// 		- a *App: Elemento tipo App
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

// Funcion Run: Inicia el servidor HTTP y habilita CORS para ejecucion local
// Receiver:
// 		- a *App: Elemento tipo App
// Input:
// 		- addr string: Direccion, para este caso el puerto
func (a *App) Run(addr string) {

	corsConfig := cors.New(cors.Options{
		AllowedHeaders: []string{"Origin", "Accept", "Content-Type", "X-Requested-With", "Authorization"},
		AllowedMethods: []string{"POST", "PUT", "GET", "PATCH", "OPTIONS", "HEAD", "DELETE"},
	})
	log.Fatal(http.ListenAndServe(addr, corsConfig.Handler(a.Router)))

	defer a.DB.Close()
}

// Funcion initializeRoutes: Define las rutas para el router
// Receiver:
// 		- a *App: Elemento tipo App
func (a *App) initializeRoutes() {
	a.Router.HandleFunc("/recetas", a.getRecetas).Methods("GET")
	a.Router.HandleFunc("/receta", a.createReceta).Methods("POST")
	a.Router.HandleFunc("/receta/{id:[0-9]+}", a.getReceta).Methods("GET")
	a.Router.HandleFunc("/receta/{id:[0-9]+}", a.updateReceta).Methods("PUT")
	a.Router.HandleFunc("/receta/{id:[0-9]+}", a.deleteReceta).Methods("DELETE")
}


// ---------- METODOS AUXILIARES ----------

// Funcion respondWithError: Procesa los errores
// Input:
// 		- w http.ResponseWriter
// 		- code int: Http status
// 		- message string: Mensaje de error
func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}

// Funcion respondWithJSON: Procesa las respuestas
// Input:
// 		- w http.ResponseWriter
// 		- code int: Http status
// 		- payload interface{}: Respuesta a enviar
func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}


// ---------- HANDLERS ----------

// Funcion getReceta: Handler para obtener la informacion de una receta
// Receiver:
// 		- a *App: Elemento tipo App
// Input:
// 		- w http.ResponseWriter
// 		- r *http.Request
func (a *App) getReceta(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Id de receta invalido")
		return
	}

	rec := Receta{IdRecetas: id}
	if err := rec.getReceta(a.DB); err != nil {

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

// Funcion getRecetas: Handler para obtener la informacion de varias recetas, filtrado por el nombre
// Receiver:
// 		- a *App: Elemento tipo App
// Input:
// 		- w http.ResponseWriter
// 		- r *http.Request
func (a *App) getRecetas(w http.ResponseWriter, r *http.Request) {
	txtBusqueda := r.FormValue("txtBusqueda")

	recetas, err := getRecetas(a.DB, txtBusqueda)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, recetas)
}

// Funcion createReceta: Handler para crear una nueva receta
// Receiver:
// 		- a *App: Elemento tipo App
// Input:
// 		- w http.ResponseWriter
// 		- r *http.Request
func (a *App) createReceta(w http.ResponseWriter, r *http.Request) {
	var rec Receta
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&rec)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Request de la carga util no valido")
		return
	}
	defer r.Body.Close()

	if err := rec.createReceta(a.DB); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusCreated, rec)
}

// Funcion updateReceta: Handler para obtener actualizar una receta
// Receiver:
// 		- a *App: Elemento tipo App
// Input:
// 		- w http.ResponseWriter
// 		- r *http.Request
func (a *App) updateReceta(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Id de receta invalido")
		return
	}

	var rec Receta
	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&rec)
	if err != nil{
		respondWithError(w, http.StatusBadRequest, "Request de la carga util no valido")
		return
	}
	defer r.Body.Close()
	rec.IdRecetas = id

	rowsAffected, err := rec.updateReceta(a.DB)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	if rowsAffected == 0 {
		respondWithError(w, http.StatusInternalServerError, "No hubo filas afectadas por la query")
		return
	}

	respondWithJSON(w, http.StatusOK, rec)
}

// Funcion deleteReceta: Handler para eliminar una receta
// Receiver:
// 		- a *App: Elemento tipo App
// Input:
// 		- w http.ResponseWriter
// 		- r *http.Request
func (a *App) deleteReceta(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Id de receta invalido")
		return
	}

	rec := Receta{IdRecetas: id}
	rowsAffected, err := rec.deleteReceta(a.DB)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	if rowsAffected == 0 {
		respondWithError(w, http.StatusInternalServerError, "No hubo filas afectadas por la query")
		return
	}

	respondWithJSON(w, http.StatusOK, map[string]string{"result": "success"})
}