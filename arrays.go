package main

import "fmt"

func main() {

	var fruits [4]string
	fmt.Println("emp:", fruits)

	fruits[0] = "Apple"
	fruits[3] = "Mango"
	fmt.Println("set:", fruits)
	fmt.Println("get:", fruits[3])

	fmt.Println("len:", len(fruits))

	scores := [5]int{10, 20, 30, 40, 50}
	fmt.Println("dcl:", scores)

	colors := [...]string{"Red", "Green", "Blue"}
	fmt.Println("dcl:", colors)

	sparse := [...]int{10, 2: 50, 90}
	fmt.Println("idx:", sparse)

	var matrix [3][2]int
	for row := range 3 {
		for col := range 2 {
			matrix[row][col] = (row + 1) * (col + 1)
		}
	}
	fmt.Println("2d: ", matrix)

	matrix = [3][2]int{
		{5, 10},
		{15, 20},
		{25, 30},
	}
	fmt.Println("2d: ", matrix)
}