package main

import "fmt"

// exemplo de como usar o init
var n int //declara fora da funcao

// FUNCAO QUE VAI SER EXECUTADA ANTES DA FUNCAO MAIN
// posso ter uma por arquivo, ñ uma por package
func init() {
	fmt.Println("Executando a função init")
	n = 10
}

func main() {
	fmt.Println("Função main sendo executada")
	fmt.Println(n)
}
