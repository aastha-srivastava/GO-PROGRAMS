package main

import "fmt"

func main() {
	var size int
	fmt.Print("Enter size of arrays: ")
	fmt.Scan(&size)

	var arr1 = make([]int, size)
	var arr2 = make([]int, size)
	var arr3 = make([]int, size)

	fmt.Println("Enter elements of first array:")
	for i := 0; i < size; i++ {
		fmt.Scan(&arr1[i])
	}

	fmt.Println("Enter elements of second array:")
	for i := 0; i < size; i++ {
		fmt.Scan(&arr2[i])
	}

	for i := 0; i < size; i++ {
		arr3[i] = arr1[i] + arr2[i]
	}

	fmt.Println("Resultant array:", arr3)
}
