package main

import "fmt"

func inverterSinal(numero int) int { //passando parametro por cópia
	return numero * -1 //passando referencia para essa funcao
}

// jogando um valor novo dentro do endereco da memoria
func inverterSinalComPonteiro(numero *int) { //parametro numero com um *PONTEIRO DE INT
	*numero = *numero * -1 //o ponteiro APONTA para um endereco de memoria
}

func main() {

	numero := 20                             //essa variável fica inalterada
	numeroInvertido := inverterSinal(numero) //cria uma nova variável, faz uma cópia do valor e manda pra funcao
	fmt.Println(numeroInvertido)
	fmt.Println(numero) //a variavel numero se mantem inalterada

	//COM PONTEIROS
	novoNumero := 40 //td alteração que fizer vai refletir na variável
	fmt.Println(novoNumero)
	inverterSinalComPonteiro(&novoNumero) //& na frente p/ referenciar onde a variavel tá salva
	fmt.Println(novoNumero)
}
