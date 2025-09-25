package main

import (
	"fmt"
	"linha-de-comando/app"
	"log"
	"os"
)

func main() {
	fmt.Println("Ponto de partida")

	aplicacao := app.Gerar()       //a aplicacao vai chamar o pct app
	erro := aplicacao.Run(os.Args) //método Run que recebe os.Args como parametro
	if erro != nil {               //o método Run retorna um erro, isso aqui é p/ tratar o erro, se houver
		log.Fatal(erro) //parecido com fmt, mas o Fatal pára o programa
	}

	//OU USANDO IF INIT
	/*if erro := aplicacao.Run(os.Args); if erro != nil {
		log.Fatal(erro)
	}*/

}
