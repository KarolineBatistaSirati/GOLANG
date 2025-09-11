package main

import "fmt"

func main() {
	fmt.Println("Maps")

	//MAPS - A CHAVE E VALOR TÊM SMP O MSM TIPO. Ñ POSSO MUDAR A ESTRUTURA.
	//DENTRO DOS COLCHETES TIPO DA CHAVE E FORA TIPO VALOR
	usuario := map[string]string{
		"nome":      "Pedro",
		"sobrenome": "Silva",
	}

	fmt.Println(usuario)
	//fmt.Println(usuario["nome"])   PEGAR SÓ INFORMACAO DO NOME,POR EX.
}
