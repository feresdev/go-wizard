package main

import "fmt"

// Declaración de constantes
const (
	nombre   = "Fernando"
	apellido = "Achocalla Estrada"
)

// Uso de iota
const (
	PrimerValor  = iota // 0
	SegundoValor        // 1
	TercerValor         // 2
)

/*
El programa imprime un mensaje de bienvenida
y muestra el nombre completo del usuario junto
con su nombre de usuario de GitHub.
*/
func main() {

	// Declaración de variables
	var usuario string = "feresdev"
	sexo := "Masculino"

	// Declaracion de constantes
	const pi float64 = 3.14159
	fmt.Println(pi)

	// Imprimir mensaje de bienvenida y nombre completo del usuario
	fmt.Println("Hola Mundo!")

	// Concatenación de cadenas
	fmt.Println("Mi nombre es: " + nombre + " " + apellido + " y mi usuario de github es " + usuario + " y mi sexo es " + sexo + ".")

	// Tipos de datos basicos
	var edad int = 22
	var altura float64 = 1.73
	var esEstudiante bool = true
	var inicial string = "F"

	fmt.Println("Edad:", edad)
	fmt.Println("Altura:", altura)
	fmt.Println("Es estudiante:", esEstudiante)
	fmt.Println("Inicial:", inicial)

	// Operadores aritmeticos (), *, /, %, +, -
	var a = (2 + 3) * 2
	fmt.Println(a)

	// Operadores de asignación: =, +=, -=, *=, /=, %=
	var b = 10
	b += 5
	fmt.Println(b)

	// Declaracion post-incremento y post-decremento: ++, --
	// (No son una expresion sino una declaracion)
	var c = 5
	c++
	fmt.Println(c)
	c--
	fmt.Println(c)

	// Operadores logicos de comparacion: ==, !=, <, >, <=, >=
	var d = 10
	var e = 20

	fmt.Println(d == e)
	fmt.Println(d != e)
	fmt.Println(d < e)
	fmt.Println(d > e)
	fmt.Println(d <= e)
	fmt.Println(d >= e)

	// Verbos de formato en fmt.Printf: %T, %v, %s
	// %T: Tipo de dato
	// %v: Valor de la variable
	// %s: Valor de la variable en formato string

	// Puntero: Variable que almacenan la direccion en memoria de un valor
	var color string = "Rojo"
	var punteroColor *string
	punteroColor = &color
	*punteroColor = "Azul"

	fmt.Printf("Tipo: %T, Valor: %s, Direccion: %v\n", color, color, &color)
	fmt.Printf("Tipo: %T, Valor: %v, Desferenciacion: %s\n", punteroColor, punteroColor, *punteroColor)

	// Arrays: Almacen un numero fijo de elementos del mismo tipo
	// var numeros [3]int
	// numeros[0] = 1
	// numeros[1] = 2
	// numeros[2] = 3

	// Declaracion de arrays con valores iniciales
	// numeros := [3]int{1, 2, 3}

	// Declaracion de arrays con valores iniciales y sin especificar el tamaño
	numeros := [...]int{1, 2, 3, 4, 5}
	fmt.Println(numeros)

	// Slice: son apuntadores a un array, no poseen datos y su tamaño puede variar
	cosas := []string{"Casa", "Carro", "Computadora"}
	autos := cosas[1:2]
	autos[0] = "Moto"
	todos := cosas[:]

	fmt.Println("Cosas:", cosas)
	fmt.Println("Autos:", autos)
	fmt.Println("Cosas[0]:", cosas[0])
	fmt.Println("Todos:", todos)

	// Tamaño y capacidad de un slice
	// len(): # de elementos en el slice
	// cap(): # de elementos de un array de origen, a partir del indice donde se creo el slice
	frutas := [5]string{"Manzana", "Pera", "Banana", "Naranja", "Uva"}
	verduras := frutas[1:3]
	verduras = append(verduras, "Lechuga", "Tomate", "Zanahoria")

	fmt.Println("Frutas:", frutas)
	fmt.Println("Verduras:", verduras)
	fmt.Println("Tamaño Verduras:", len(verduras))
	fmt.Println("Capacidad Verduras:", cap(verduras))

	// Slice usando make() cuando no se conoce el tamaño del array
	estudiantes := make([]string, 0, 3)
	estudiantes = append(estudiantes, "Juan", "Maria", "Pedro")

	fmt.Println("Estudiantes:", estudiantes)
	fmt.Println("Tamaño Estudiantes:", len(estudiantes))
	fmt.Println("Capacidad Estudiantes:", cap(estudiantes))

	// Nota: Si se supera la capacidad del slice se duplicara la capacidad del slice

	// Map: Estructura de datos que almacena pares clave-valor
	musicas := make(map[string]string)
	musicas["Chamame"] = "Me ire de ti amooor juro me ireeee"
	musicas["Pop"] = "Set Fire to The Rain"

	fmt.Println(musicas)

	// Mapas literales
	musicas2 := map[string]string{
		"Chamame": "Me contaroooon que tu noovio es un ricoooo !!!",
		"Pop":     "Rolling in The Deep",
	}

	fmt.Println(musicas2)

	delete(musicas2, "Pop")
	fmt.Println(musicas2)

	fmt.Println(musicas2["Chamame"])

	// Valores que devuelve un mapa
	content, ok := musicas2["Chamame"]
	fmt.Println(content, ok)

	// Estructuras: Permiten almacenar coleccion de campos, algo muy parecido a las clases de otros lenguajes
	type Persona struct {
		Nombre     string
		Edad       int
		TieneHijos bool
	}

	fernando := Persona{
		Nombre:     "Fernando",
		Edad:       22,
		TieneHijos: false,
	}

	fmt.Printf("%+v", fernando)
}
