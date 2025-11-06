package main

import (
	"fmt"
)

func RemoveElements[T comparable](lista []T, aBorrar []T) []T {
	del := make(map[T]struct{}, len(aBorrar))
	for _, v := range aBorrar {
		del[v] = struct{}{}
	}
	keep := func(x T) bool { // lambda/predicado
		_, found := del[x]
		return !found
	}
	out := make([]T, 0, len(lista))
	for _, x := range lista {
		if keep(x) {
			out = append(out, x)
		}
	}
	return out
}

func main() {
	original := []string{"rojo", "verde", "azul", "amarillo", "gris", "blanco", "negro"}
	aBorrar := []string{"amarillo", "café", "blanco"}

	res := RemoveElements(original, aBorrar)
	fmt.Println(res) // [rojo verde azul gris negro]
}
