package main

import (
	"fmt"
	"modulo/auxiliar"

	"github.com/badoux/checkmail" //importei um pacote externo
)

func main() {
	fmt.Println("Escrevendo do arquivo main")
	auxiliar.Escrever()

	erro := checkmail.ValidateFormat("estudosgo@gmail.com") //chamando a função de um pacote externo
	fmt.Println(erro)
}
