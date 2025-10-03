package main

import "fmt"

func main() {
	var size int
	fmt.Print("Enter size of slices: ")
	fmt.Scan(&size)

	slice1 := make([]int, size)
	slice2 := make([]int, size)
	slice3 := make([]int, size)

	fmt.Println("Enter elements of first slice:")
	for i := 0; i < size; i++ {
		fmt.Scan(&slice1[i])
	}

	fmt.Println("Enter elements of second slice:")
	for i := 0; i < size; i++ {
		fmt.Scan(&slice2[i])
	}

	for i := 0; i < size; i++ {
		slice3[i] = slice1[i] + slice2[i]
	}

	fmt.Println("Resultant slice:", slice3)
}
