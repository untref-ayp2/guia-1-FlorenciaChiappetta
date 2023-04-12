package busquedas

import (
	"reflect"
	"testing"
)

func TestOrdenDescendente(t *testing.T) {

	arreglo := []int{90, -2, 7, -4, 5, 88, 40, 20}
	desired := []int{90, 88, 40, 20, 7, 5, -2, -4}
	result := BurbujeoDescendente(arreglo)

	if !reflect.DeepEqual(result, desired) {
		t.Errorf("Funcion de ordenamiento por burbujeo = %v, deseado= %v/n", result, desired)
	}
}
func TestOrdenAscendente(t *testing.T) {

	arreglo := []int{90, -2, 7, -4, 5, 88, 40, 20}
	desired := []int{-4, -2, 5, 7, 20, 40, 88, 90}

	result := BurbujeoAscendente(arreglo)
	if !reflect.DeepEqual(result, desired) {
		t.Errorf("Funcion de ordenamiento por burbujeo = %v  deseado= %v\n", result, desired)
	}
}

func TestOrdenDescendente2(t *testing.T) {

	arreglo := []int{90, -2, 7, -4, 5, 88, 40, 20}
	desired := []int{90, 88, 40, 20, 7, 5, -2, -4}
	result := BurbujeoAscendente(arreglo)

	if !reflect.DeepEqual(result, desired) {
		t.Errorf("Funcion de ordenamiento por burbujeo = %v, deseado= %v/n", result, desired)
	}
}
func TestOrdenAscendente2(t *testing.T) {

	arreglo := []int{90, -2, 7, -4, 5, 88, 40, 20}
	desired := []int{-4, -2, 5, 7, 20, 40, 88, 90}

	result := BurbujeoDescendente(arreglo)
	if !reflect.DeepEqual(result, desired) {
		t.Errorf("Funcion de ordenamiento por burbujeo = %v  deseado= %v\n", result, desired)
	}
}
