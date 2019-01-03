package main

// Funcion main: Punto de entrada a la aplicacion
func main() {
	a := App{}

	a.Initialize()

	a.Run(":3000")
}
