package main

import (
	"fmt"
	"sync"
	"time"
)

// waitgroup é
func main() {
	//crio uma variavel
	var waitGroup sync.WaitGroup //package sync e tem esse tipo waitgroup

	waitGroup.Add(2) //2 - qntd de goroutine rodando. Qntd de goroutine no grupodeespera

	go func() { //cria uma funcao goroutine
		escrever("Olá, mundo!") //chama a funcao escrever
		waitGroup.Done()        //tira uma do contador, acima contém 2.
	}()

	go func() {
		escrever("Programando em Go!")
		waitGroup.Done()
	}()

	waitGroup.Wait() //espera a contagem das goroutines chegar em 0
}

func escrever(texto string) {
	//looping com parametros estabelecidos
	for i := 0; i < 5; i++ {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}
