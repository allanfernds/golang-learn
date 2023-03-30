package main

import (
	"fmt"
)

func main() {
	var number int = 23
	fmt.Println(parouimpar(number))
}

func parouimpar(number int) string {
	if (number % 2) == 0 {
		return "Par"
	} else {
		return "Impar"
	}

}
