package main

import (
	"busquedas"
	"fmt"
)

func main() {
	/*arreglo := utiles.GenerarArreglo(10, 100000000)
	buscado := -1

	//fmt.Println(arreglo)

	inicio := time.Now()
	// Busqueda Lineal
	busquedas.BusLineal(arreglo, buscado)
	fmt.Println("Busqueda Lineal: ", time.Since(inicio))

	inicio = time.Now()
	// Ordenar el arreglo para la busqueda binaria
	sort.Ints(arreglo)
	fmt.Println("Ordenamiento: ", time.Since(inicio))
	//fmt.Println(arreglo)

	inicio = time.Now()
	// Busqueda Binaria
	busquedas.BusquedaBinaria(arreglo, buscado)
	fmt.Println("Busqueda Binaria: ", time.Since(inicio))

	/*Burbujeo(arreglo)
	fmt.Println("Burbujeo: ", time.Since(inicio))

	Selection(arreglo)
	fmt.Println("Seleccion: ", time.Since(inicio))

	Insertion(arreglo)
	fmt.Println("Insercion: ", time.Since(inicio))
	// ordenamiento por Burbuja, Seleccion e Inserción
	/*numbers := []int{9, -2, 22, 3, 7, 2, 60, 71, 8}
	items := []int{9, -2, 22, 3, 7, 2, 60, 71, 8}
	prueba := []int{2, 1, -1, 4, 3}
	Burbujeo(numbers)
	fmt.Println(numbers)
	Selection(items)
	fmt.Println(items)
	Insertion(prueba)
	fmt.Println(prueba)*/

	arreglo2 := []int{90, -2, 7, -4, 5, 88, 40, 20}
	b := busquedas.BurbujeoAscendente(arreglo2)
	fmt.Println(b)

}
