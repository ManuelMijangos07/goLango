package main

// import "time"

//"time"
// estructura de datos
/* type Gorra struct {
	marca  string
	color  string
	precio float32
	plana  bool
} */

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

	//estructura de datos llena
	/* var gorraNegra = Gorra{
		marca:  "Nike",
		color:  "Negro",
		precio: 25.50,
		plana:  false}
	fmt.Print(gorraNegra) */

	//llamada de funciones
	/* fmt.Println(gorras(45, "EUROS"))
	pantalon("rojo", "largo", "sin bolsillos", "nike") */

	// DEFINIR UN ARRAY FORMA 1
	/* var peliculas [3]string
	peliculas[0] = "La verdad duele"
	peliculas[1] = "Ciudadano Ejemplar"
	peliculas[2] = "Gran torino"

	fmt.Println(peliculas[1]) */

	//declarar arrays forma 2
	/* peliculas := [3]string{
		"la verdad duele",
		"ciudadano ejemplar",
		"batman"}
	fmt.Println(peliculas[1]) */

	//Arrays multidimencional
	/* var peliculas [3][2]string
	peliculas[0][0] = "La verdad duele"
	peliculas[0][1] = "Mientras duermes"
	peliculas[1][0] = "Ciudadno enjemplar"
	peliculas[1][1] = "El señor de los anillos"
	peliculas[2][0] = "Gran Torino"
	peliculas[2][1] = "A todo gas"

	fmt.Println(peliculas) */

	//slice
	/* peliculas := []string{
		"La verdad duele",
		"ciudadano ejemplar"}
	fmt.Println(peliculas) */

	//agregar elementos al array
	//peliculas = append(peliculas, "sin límites")

	// longitud de array
	//fmt.Println(len(peliculas))
	// imprimir elementos de array desde hasta
	//fmt.Println(peliculas[0:3])
}

//func con closures
/* func gorras(pedido float32, moneda string) (string, float32, string) {
	precio := func() float32 {
		return pedido * 7
	}
	return "numero de gorras pedidas:", precio(), moneda
} */

//func mult atributos
/* func pantalon(atributos ...string) {
	for _, atributo := range atributos {
		fmt.Println(atributo)
	}
}
*/
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
