package main

import "fmt"

func main() {
	fmt.Println("Ponteiros")

	//PONTEIRO uma variável que salva o endereço de memória de alguma coisa
	//PONTEIRO É UM REFERÊNCIA DE MEMÓRIA
	variavel1 := 10
	variavel2 := variavel1

	fmt.Println(variavel1, variavel2)

	variavel1++ //isso altera só a variavel1 e pq a 2 é uma cópia

	//DIFERENÇA ENTRE ATRIBUIR VALOR ASSIM OU COM PONTEIROS?
	var variavel3 int //o valor inicial é 0 e guarda, literalmente, o valor int
	var ponteiro *int //o valor inicial é <NIL> e guardar o enreço de memória de um valor int

	variavel3 = 100
	ponteiro = &variavel3 //na frente da variável coloque um &

	fmt.Println(variavel3, ponteiro)
}
