package main

import (
	"fmt"
	"time"
)

// CONCORRÊNCIA(revezam o tempo de execução das tarefas, pode ser duas ou +)!= PARALELISMO(2 ou mais tarefas executadas ao msm tempo)
// Goroutine é uma funcao/metodo q é chamado quando coloca o go
func main() {
	go escrever("Olá, mundo!") //goroutine - roda ao msm tempo, mas de forma alternada
	escrever("Programando em Go!")
}

func escrever(texto string) {
	//looping infinito
	for {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}
