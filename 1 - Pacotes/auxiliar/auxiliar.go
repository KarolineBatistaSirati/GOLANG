package auxiliar

import "fmt"

// Escrever registra uma mensagem na tela
func Escrever() { //p/ exportar uma funcao, ela precisa estar com a inicial maiuscula
	fmt.Println("Escrevendo do pacote auxiliar")
	escrever2() //dentro do mesmo pacote consigo chamar as funcoes com letra minuscula, fora dele não.
}

//sempre que exportar uma funcao, escreva um comentario acima dela dizendo o que faz
