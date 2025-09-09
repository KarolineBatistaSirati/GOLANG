package main

import (
	"errors"
	"fmt"
)

func main() {
	//TIPOS DE NUMEROS INTEIROS
	//int8, int16, int32, int64   A diferença é o tamanho deles e aceitam negativos
	var numero int64 = 10000000000000
	fmt.Println(numero)

	//mesma coisa do de cima, mas aqui é por inferência
	numero2 := 10000000000000
	fmt.Println(numero2)

	//uint = unsigned int (não permite numeros negativos)
	var numero3 uint32 = 10000
	fmt.Println(numero3)

	//alias
	//INT32 = RUNE
	var numero4 rune = 1234
	fmt.Println(numero4)

	//BYTE = UINT8
	var numero5 byte = 123
	fmt.Println(numero5)
	//FIM NUMEROS INTEIROS

	//FLOAT
	//float32, float 64
	var numeroReal1 float32 = 123.45
	fmt.Println(numeroReal1)
	//msm coisa
	numeroReal2 := 12345.67
	fmt.Println(numeroReal2)
	//FIM NUMEROS REAIS

	//STRINGS
	var str string = "Texto"
	fmt.Println(str)
	//mesma coisa
	str2 := "Texto2"
	fmt.Println(str2)
	//FIM STRINGS

	//VALOR ZERO
	var texto int16
	fmt.Println(texto)
	//pra string é vazio e pra int é zero, pq td tipo de dado tem um valor inicial
	texto2 := 5
	fmt.Println(texto2) // é o msm que o de cima, mas aqui é obrigatório colocar um valor

	//BOOLEANO
	var booleano1 bool = true //O valor inicial é FALSE se não atribuir valor
	fmt.Println(booleano1)

	//ERRO
	var erro error = errors.New("Erro interno") // errors é um pacote dentro do GO
	fmt.Println(erro)                           // o valor inicial é NIL = zero
}
