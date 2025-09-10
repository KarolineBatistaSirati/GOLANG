package main

import "fmt"

type usuario struct { //type (pra criar um tipo) + nome do tipo + estrutura p/ aguardar mts campos
	nome     string
	idade    uint8
	endereco endereco //struct dentro de struct
}

type endereco struct {
	logradouro string
	numero     uint8
}

//STRUCTS é uma colação de campos, cada um tem seu nome e tipo
//Usa ele quando se tem várias informações relacionadas sobre a msm coisa

func main() {
	fmt.Println("Arquivo structs")

	// primeiro jeito de declarar o struct
	var u usuario
	u.nome = "Karol"
	u.idade = 28
	fmt.Println(u)

	enderecoExemplo := endereco{"Rua dos bobos", 0}

	// segundo jeito de declarar o struct, por inferência, é o + rápido
	usuario2 := usuario{"Karol", 28, enderecoExemplo}
	fmt.Println(usuario2)

	// terceiro jeito, quando ñ tenho tds as info do usuário
	usuario3 := usuario{nome: "Karol"} //preenche o primeiro campo e o segundo fica zerado
	fmt.Println(usuario3)
}
