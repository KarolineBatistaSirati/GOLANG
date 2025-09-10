package main

import "fmt"

//FUNÇÃO É UM SÉRIE DE PASSOS QUE SERÃO SEGUIDOS PELO PROGRAMA
//Elas podem ter parâmetros (dado que coloca na funcao p/ funcionar) e retornos (o que a funcao devolve)

// as funções podem ter vários retornos
func calculosMatematicos(n1, n2 int8) (int8, int8) { //parametros + retornOS
	soma := n1 + n2
	subtracao := n1 - n2
	return soma, subtracao
}

func somar(n1 int8, n2 int8) int8 { //palavra chave + parâmetro + retorno
	return n1 + n2
}

func main() {
	soma := somar(10, 20)
	fmt.Println(soma)

	//declarar uma variável do tipo função e jogar uma função dentro dela
	var f = func(txt string) string {
		fmt.Println(txt) //além de printar, vai retornar.
		return txt
	}

	resultado := f("Texto da função 1")
	fmt.Println(resultado)

	resultadoSoma, ResultadoSubtracao := calculosMatematicos(10, 15) //Se criou a funcao, mas ñ quer usar, coloque o _ EXEMPLO: resultadoSoma, _ := calculosMatematicos(10, 15)
	fmt.Println(resultadoSoma, ResultadoSubtracao)
}
