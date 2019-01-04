# Recetas (Golang rest api)

> Rest api desarrollada en Golang para  realizar el CRUD a datos de recetas de cocina y buscar por su nombre. Implementada utilizando con Gorilla Mux y base de datos Cockroachdb.

## Funciones CRUD para Recetas

- Crear: POST en /receta
- Actualizar: PUT en /receta/{id}
- Eliminar: DELETE en /receta/{id}
- Obtener una: GET en /receta/{id}
- Obtener todas: GET en /recetas con parámetro 'txtBusqueda' (Texto para buscar por nombre)

## Estructura de datos

Receta:
- idrecetas int
- nombre string
- tipoplato string
- porciones int
- ingredientes string
- preparacion string

## URLs

- Rest Api: http://localhost:3000
- Base de datos Cockroachdb: http://localhost:26257
