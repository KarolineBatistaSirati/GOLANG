package main

import "fmt"

func fibonacci(posicao uint) uint {
	if posicao <= 1 {
		return posicao
	}

	return fibonacci(posicao-2) + fibonacci(posicao-1)
}

func main() {
	fmt.Println("Recursivas")
	//RECURSIVA = p/ funcionar ela depende de outra excecução dela msm

	posicao := uint(10)
	fmt.Println(fibonacci(posicao))
}
