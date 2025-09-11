package main

import "fmt"

func main() {
	fmt.Println("Estruturas de controle")

	//AVALIÇÃO DE UMA CONDIÇÃO. SE FOR TRUE, FAZ ALGO; SE FOR FALSE, FAZ OUTRA COISA.

	numero := -5
	if numero > 15 {
		fmt.Println("O número é maior que 15")
	} else {
		fmt.Println("O número é menor ou igual a 15")
	}

	//IF INIT
	/*if outroNumero := numero; outroNumero > 0 {
		fmt.Println("Número é maior que zero")
	} else {
		fmt.Println("Número é menor que zero")
	}*/

	if outroNumero := numero; outroNumero > 0 {
		fmt.Println("Número é maior que zero")
	} else if numero < -10 {
		fmt.Println("Número é menor que -10")
	} else {
		fmt.Println("Entre 0 e -10")
	}

}
