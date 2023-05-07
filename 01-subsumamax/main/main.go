package main

import (
	"fmt"
	"subsumamax"
)

func main() {
	//arreglo := []int{-1, 6, -2, 5, -1, 4, 3, -4, 3}
	arreglo := []int{-2, 11, -4, 2, -3, -10}
	//sumaMaxima, posInicial, posFinal := subsumamax.SubsecuenciaSumaMaxima(arreglo)
	//fmt.Printf("La subsecuencia con suma máxima es: %v\n", arreglo[posInicial:posFinal+1])
	//fmt.Printf("Punto Critico en función: %v\n", sumaMaxima)
	//fmt.Printf("La posición inicial es: %v\n", posInicial)
	//fmt.Printf("La posición final es: %v\n", posFinal)
	sumaMaxima, sumaDerechaAIzquierda, sumaIzquierdaADerecha, sumaCentral := subsumamax.SubsecuenciaSumaMaxima2(arreglo)
	fmt.Printf("La suma máxima es: %v\n,%v\n,%v\n,%v\n", sumaMaxima, sumaDerechaAIzquierda, sumaIzquierdaADerecha, sumaCentral)
	fmt.Printf("   *   \n")
	fmt.Printf("  * *  \n")
	fmt.Printf(" * * * \n")
	fmt.Printf("* * * * \n")

	b := "PRUEBA"
	ASCII := []byte(b)

	fmt.Println(ASCII)
}
