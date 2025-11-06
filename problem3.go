package main

import (
	"errors"
	"fmt"
)

func Transpose[T any](X [][]T) ([][]T, error) {
	if len(X) == 0 {
		return [][]T{}, nil
	}
	m, n := len(X), len(X[0])

	// Validar que la matriz sea rectangular
	for i := 1; i < m; i++ {
		if len(X[i]) != n {
			return nil, errors.New("matriz irregular: todas las filas deben tener la misma longitud")
		}
	}

	// Lambda que construye la fila j-ésima de X^T a partir de la columna j de X
	makeRow := func(j int) []T {
		row := make([]T, m)
		for i := 0; i < m; i++ {
			row[i] = X[i][j]
		}
		return row
	}

	XT := make([][]T, n)
	for j := 0; j < n; j++ {
		XT[j] = makeRow(j)
	}
	return XT, nil
}

func main() {
	X := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	XT, err := Transpose(X)
	if err != nil {
		panic(err)
	}
	for _, r := range XT {
		fmt.Println(r)
	}
}
