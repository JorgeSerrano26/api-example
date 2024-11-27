package Structures

import "fmt"

func matrices() {
	var a [5]int // Declaración de un array de 5 elementos de tipo int

	var b = [5]int{1, 2, 3, 4} // Declaración de un array de 5 elementos de tipo int con valores iniciales

	c := [5]int{1, 2, 3, 4, 5} // Declaración de un array de 5 elementos de tipo int con valores iniciales

	fmt.Println(a, b, c)

	// Recorrer un array
	for i := 0; i < len(c); i++ {
		fmt.Println(c[i])
	}

	// Declaración de una matriz de 3x3
	var matrix [3][3]int

	matrix[0] = [3]int{1, 2, 3}
	matrix[1] = [3]int{4, 5, 6}
	matrix[2] = [3]int{7, 8, 9}

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			fmt.Println(matrix[i][j])
		}
	}

	for index, value := range matrix {
		for index2, value2 := range value {
			fmt.Printf("Row %d, Column, Column %d, value %d\n", index, index2, value2)
		}
	}
}
