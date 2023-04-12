package subsumamax

// Dado un arreglo devuelve la posicion inicial, la posición final y el valor
// de la subsecuencia cuya suma es máxima dentro del arreglo original
func SubsecuenciaSumaMaxima(arreglo []int) (int, int, int) {
	sumaMaxima := 0
	posInicial, posFinal := -1, -1
	for i := 0; i < len(arreglo); i++ {
		sumaLocal := 0
		for j := i; j < len(arreglo); j++ {
			sumaLocal += arreglo[j]
			if sumaLocal > sumaMaxima {
				sumaMaxima = sumaLocal
				posInicial = i
				posFinal = j
			}
		}
	}
	return sumaMaxima, posInicial, posFinal
}

func SubsecuenciaSumaMaxima2(arreglo []int) (int, int, int, int) {
	sumaMaxima := 0
	k := len(arreglo) - 1
	sumaDerechaAIzquierda := 0
	sumaIzquierdaADerecha := 0
	sumaCentral := 0
	for i := 0; i < len(arreglo); i++ {
		sumaMaxima += arreglo[i]
		sumaDerechaAIzquierda = sumaMaxima - arreglo[k]
		sumaIzquierdaADerecha = sumaMaxima - arreglo[i]
		sumaCentral = sumaMaxima - (-(arreglo[i])) - (-(arreglo[k]))
		k--
	}

	if sumaMaxima < sumaDerechaAIzquierda {
		sumaMaxima = sumaDerechaAIzquierda
	}
	if sumaMaxima < sumaIzquierdaADerecha {
		sumaMaxima = sumaIzquierdaADerecha
	}
	if sumaMaxima < sumaCentral {
		sumaMaxima = sumaCentral
	}
	return sumaMaxima, sumaDerechaAIzquierda, sumaIzquierdaADerecha, sumaCentral
}
