package busquedas

func BurbujeoAscendente(ListaDesordenada []int) []int {
	var auxiliar int
	for i := 0; i < len(ListaDesordenada); i++ {
		for j := 0; j < len(ListaDesordenada); j++ {
			if ListaDesordenada[i] < ListaDesordenada[j] {
				auxiliar = ListaDesordenada[i]
				ListaDesordenada[i] = ListaDesordenada[j]
				ListaDesordenada[j] = auxiliar
			}
		}
	}
	return ListaDesordenada
}

func BurbujeoDescendente(ListaDesordenada []int) []int {
	var auxiliar int
	for i := 0; i < len(ListaDesordenada); i++ {
		for j := 0; j < len(ListaDesordenada); j++ {
			if ListaDesordenada[i] > ListaDesordenada[j] {
				auxiliar = ListaDesordenada[i]
				ListaDesordenada[i] = ListaDesordenada[j]
				ListaDesordenada[j] = auxiliar
			}
		}
	}
	return ListaDesordenada
}
