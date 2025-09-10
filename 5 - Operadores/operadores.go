package main

import "fmt"

func main() {
	//ARITMÉTICOS (CONTAS)
	soma := 1 + 2
	subtracao := 2 - 1
	divisao := 10 / 4
	multiplicacao := 10 * 5
	restoDaDivisao := 10 % 2

	fmt.Println(soma, subtracao, divisao, multiplicacao, restoDaDivisao)

	//ñ posso fazer nada com dados diferente, neste caso:
	/*var numero1 int16 = 10
	var numero2 int32 = 25
	soma2 := numero1 + numero2 //aqui só daria certo se fossem do msm tipo, int16 int16 ou 32 nos dois.
	fmt.Println(soma)*/
	//FIM DOS ATIRMÉTICOS

	//ATRIBUIÇÃO
	var variavel1 string = "String"
	variavel2 := "String2"
	fmt.Println(variavel1, variavel2)
	//FIM DOS OPERADORES DE ATRIBUIÇÃO

	//OPERADORES RELACIONAIS
	fmt.Println(1 > 2)
	fmt.Println(1 >= 2)
	fmt.Println(1 == 2)
	fmt.Println(1 <= 2)
	fmt.Println(1 > 2)
	fmt.Println(1 != 2) //diferente
	//FIM OPERADORES RELACIONAIS

	//OPERADORES LÓGICOS
	fmt.Println("-----------------")
	verdadeiro, falso := true, false
	fmt.Println(verdadeiro && falso) //E
	fmt.Println(verdadeiro || falso) //OU
	fmt.Println(!verdadeiro)         //NEGAÇÃO (aparece o valor contrário do que está escrito)
	fmt.Println(!falso)
	//FIM DOS OPERADORES LÓGICOS

	//OPERADORES UNÁRIOS (interagem com uma variável por vez)
	numero := 10
	numero++     //aumentar o valor em 1
	numero += 15 //aumentar mais de 1
	fmt.Println(numero)

	//decrementar
	numero--
	numero -= 20
	fmt.Println(numero)
	//FIM DOS OPERADORES UNÁRIOS

	//OPERADORES TERNÁRIOS
	var texto string
	if numero > 5 {
		texto = "Maior que 5"
	} else {
		texto = "Menor que 5"
	}

	fmt.Println(texto)

}
