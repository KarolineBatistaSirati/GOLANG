package main

import "fmt"

func main() {

	//COMO EXECUTAR UMA FUNÇÃO QUE Ñ TEM NOME: com () no final
	/*func() {
		fmt.Println("Olá, mundo!")
	}()*/

	//SE QUISER PASSAR PARAMETROS
	/*func(texto string) {
		fmt.Println(texto)
	}("Passando parâmetro")*/ //aqui coloca o valor do parametro

	//ao invés de imprimir, retornar
	retorno := func(texto string) string {
		return fmt.Sprintf("Recebido -> %s", texto)
	}("Passando parâmetro") //aqui coloca o valor do parametro

	fmt.Println(retorno)

	//Sprintf = passar um string dentro e concatena com outros tipos de dados

}
