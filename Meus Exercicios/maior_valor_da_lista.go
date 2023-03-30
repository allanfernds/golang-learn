package main

import "fmt"

func main() {
	fmt.Println(maiorValor([]int{10, 23, 1, 23, 45, 67, 3, 4}))
}

func maiorValor(lista []int) int {
	var count int = 0

	for i := 0; i < len(lista); i++ {
		if lista[i] > count {
			count = lista[i]
		}
	}

	return count
}
