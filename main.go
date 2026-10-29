package main

import "fmt"

// Declaración de constantes
const (
	nombre   = "Fernando"
	apellido = "Achocalla Estrada"
)

/*
El programa imprime un mensaje de bienvenida
y muestra el nombre completo del usuario junto
con su nombre de usuario de GitHub.
*/
func main() {
	var usuario string = "feresdev"
	fmt.Println("Hola Mundo!")
	fmt.Println("Mi nombre es: " + nombre + " " + apellido + " y mi usuario de github es " + usuario + ".")
}
