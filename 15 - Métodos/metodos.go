package main

import "fmt"

//Normalmente, p/ chamar funcoes
/*func escrever() {
	fmt.Println("Escrevendo...")
}

func main() {
	escrever()
}*/
//MÉTODO X FUNCAO
//MÉTODO TÁ LIGADO A ALGUMA ESTRUTURA,

//COM MÉTODOS

type usuario struct {
	nome  string
	idade uint8
}

func (u usuario) salvar() { //é como se fosse uma funcao, mas é um método nomeado salvar(), e o usuario é a estrutura do método
	fmt.Println("Dentro do método salvar") //esse usuario é porque precisa ter o msm nome do type
}

// MÉTODO COM RETORNO
func (u usuario) maiorDeIdade() bool { //ESSA ALTERACAO ACONTECE SÓ DENTRO DO MÉTODO
	return u.idade >= 18
}

// COM PONTEIRO							//alteracao na referencia de memoria
func (u *usuario) fazerAniversario() { //como o usuario é um ponteiro, fora tbm sofre alteracao
	u.idade++
}

func main() {
	usuario1 := usuario{"Usuário 1", 20}
	usuario1.salvar() //p/ chamar o método

	usuario2 := usuario{"Davi", 30}
	usuario2.salvar()

	//ACESSANDO
	maiorDeIdade := usuario2.maiorDeIdade()
	fmt.Println(maiorDeIdade)

	usuario2.fazerAniversario()
	fmt.Println(usuario2.idade)
}
