package main

import "fmt"

func funcao1() {
	fmt.Println("Executando a função 1")
}

func funcao2() {
	fmt.Println("Executando a função 2")
}

// DEFER COM RETORNO
func alunoEstaAprovado(n1, n2 float32) bool {
	defer fmt.Println("Média calculada. Resultado será retornado!") //ELA APARECE ANTES DO RETORNO FINAL
	fmt.Println("Entrando na função para ver se o aluno está aprovado")

	media := (n1 + n2) / 2

	if media >= 6 {
		return true
	}
	return false
}

func main() {

	//PARA CHAMAR AS FUNCOES ACIMA
	/*funcao1()
	funcao2()*/

	//USANDO O DEFER (ADIAR)
	/*defer funcao1()
	funcao2()*/

	fmt.Println(alunoEstaAprovado(7, 8))

}
