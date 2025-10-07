package main

import "fmt"

func main() {
	//COM BUFFER SÓ BLOQUEIA QNDO ATINGIR A CAPACIDADE MÁX. DELE, NESTE CASO: 2canais
	canal := make(chan string, 2) //criando canal de strings
	canal <- "Olá Mundo!"         //enviando o valor olá mundo p/ canal
	//aqui não posso colocar mais canais por conta da capacidade que estabeleci acima 2

	mensagem := <-canal //variavel mensagem que vai receber o canal
	mensagem2 := <-canal

	fmt.Println(mensagem) //printo a variavel
	fmt.Println(mensagem2)
}
