package main

import "fmt"

func main() {
	var num int
	fmt.Print("Enter a number: ")
	fmt.Scan(&num)

	isPrime := true
	if num < 2 {
		isPrime = false
	}
	for i := 2; i*i <= num; i++ {
		if num%i == 0 {
			isPrime = false
			break
		}
	}

	if isPrime {
		fmt.Println(num, "is Prime")
	} else {
		fmt.Println(num, "is Not Prime")
	}
}
