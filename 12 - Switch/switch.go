package main

import "fmt"

// SWITCH PRA TESTAR MÚLTIPLAS CONDIÇÕES
func diaDaSemana(numero int) string {
	switch numero { //avalia essa variável
	case 1: //se caso for igual a 1
		return "Domingo" //faça isso, retorne domingo
	case 2: //se caso for igual a 2
		return "Segunda-feira" //faça isso, retorne segunda-feira
	case 3:
		return "Terça-feira"
	case 4:
		return "Quarta-feira"
	case 5:
		return "Quinta-feira"
	case 6:
		return "Sexta-feira"
	case 7:
		return "Sábado"
	default: //vai acontecer se nenhuma das cond acima for atendida
		return "Número inválido"
	}
}

// outra maneira:
func diaDaSemana2(numero int) string {
	switch {
	case numero == 1:
		return "Domingo"
	default: //vai acontecer se nenhuma das cond acima for atendida
		return "Número inválido"
	}
}

func main() {
	fmt.Println("Switch")

	dia := diaDaSemana(200)
	fmt.Println(dia)
}
