package main

import "fmt"

// import "time"

//"time"

type Gorra struct {
	marca  string
	color  string
	precio float32
	plana  bool
}

func main() {
	//time.Sleep(time.Second * 5) // retardo al ejecutar el programa

	// declaración de variables
	// var suma int = 8 + 9
	// var resta int = 6 - 4
	// var apelliedos string = "Mijangos Campos"

	// declaración segunda forma
	// pais := "España"
	// var prueba bool = true
	// fmt.Print(prueba)

	//tipos de datos
	/*
		float32
		float64
		int
		bool
		string
	*/

	// //constantes
	// const year int = 2018

	// llamada de funciones
	// prueba();

	var gorraNegra = Gorra{
		marca:  "Nike",
		color:  "Negro",
		precio: 25.50,
		plana:  false}

	fmt.Print(gorraNegra)

	fmt.Println(gorras(45, "EUROS"))

}

func gorras(pedido float32, moneda string) (string, float32, string) {
	precio := func() float32 {
		return pedido * 7
	}
	return "numero de gorras pedidas:", precio(), moneda
}

// definir funciones
/* func prueba(){

}*/

// funcion con datos.
/*
 func operacion(n1 float32, n2 float32, op string) {
	...
 } */

//funcion con retorno de objetos segun tipo y cantidad
/* func test() (string, string){

}
*/

// Closures
