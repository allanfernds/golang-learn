package main

import "fmt"

func main() {
	var numbers = []int{1, 8, 2, 3}
	fmt.Println(somaDePares(numbers))
}

func somaDePares(numeros []int) int {
	var count int = 0
	for i := 0; i < len(numeros); i++ {
		if (numeros[i] % 2) == 0 {
			count += numeros[i]
		}
	}
	return count
}
