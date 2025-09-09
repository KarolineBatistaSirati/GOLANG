package main

import "fmt"

func main() {
	//DECLARANDO VARIÁVEIS
	var variavel1 string = "Variável 1" //deixando o tipo da var explícito
	variavel2 := "Variável 2"           //implícito e mais comum. Ele determina o tipo pelo valor.
	fmt.Println(variavel1)
	fmt.Println(variavel2)

	//DECLARANDO VÁRIAS VARIÁVEIS DE UMA VEZ
	var (
		variavel3 string = "Variável 3"
		variavel4 string = "Variável 4"
	)
	fmt.Println(variavel3, variavel4)

	//Ouu assim
	variavel5, variavel6 := "Variável 5", "Variável 6"
	fmt.Println(variavel5, variavel6)

	//CONSTANTES
	const constante1 string = "Constante 1"
	fmt.Println(constante1)

	//INVERTER VARIAVEIS
	variavel5, variavel6 = variavel6, variavel5
	fmt.Println(variavel5, variavel6)
}
