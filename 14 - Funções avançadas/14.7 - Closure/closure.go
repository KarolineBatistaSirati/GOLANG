package main

import "fmt"

// FUNÇÕES QUE REFERENCIAM VARIÁVEIS FORA DO SEU CORPO
func closure() func() {
	texto := "Dentro da função closure"

	funcao := func() {
		fmt.Println(texto)
	}

	return funcao

}

func main() {
	//antes de chamar a func closure dentro da func main
	texto := "Dentro da função main"
	fmt.Println(texto)

	funcaoNova := closure()
	funcaoNova()

	// closure = é como se tivesse uma "memória" de onde ela veio, quando der println, sempre será chamando o texto de quando foi declarada inicialmente
}
