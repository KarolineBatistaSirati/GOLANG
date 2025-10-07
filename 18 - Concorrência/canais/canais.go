package main

import (
	"fmt"
	"time"
)

func main() {
	//como criar um canal?
	//use a func make, chan é palavra chave
	//esse só pode enviar e receber dados do tipo string
	canal := make(chan string)
	go escrever("Olá mundo", canal)

	fmt.Println("Depois da função escrever começar a ser executada")

	for {
		mensagem, aberto := <-canal //variaveis que recebem o valor (<-canal = recebe valor); aberto (ver se o canal tá aberto)
		if !aberto {
			break //chamar o comando break é uma forma de sair do looping infinito
		}

		fmt.Println(mensagem)
	}

	//REFATORANDO O FOR
	/*for mensagem := range canal {
		fmt.Println(mensagem)
	}
	*/

	fmt.Println("Fim do programa!")
}

// channel é um canal de comunicacao, envia e recebe dados
func escrever(texto string, canal chan string) {
	for i := 0; i < 5; i++ {
		canal <- texto //o canal vai receber este texto (canal<- é pra enviar valor pra dentro do canal)
		time.Sleep(time.Second)
	}

	close(canal) //ao finalizar o looping, feche o canal
}
