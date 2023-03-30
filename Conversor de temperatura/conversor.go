package main

import "fmt"

func main() {
	fmt.Println("Digite a temperatura em graus celsius")
	conversor()
}

func conversor() float64 {
	var temp float64
	fmt.Scanln(&temp)
	fmt.Println(&temp)
	return (temp * 1.8) + 32
}
