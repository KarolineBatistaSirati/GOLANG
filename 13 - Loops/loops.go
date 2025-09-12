package main

import (
	"fmt"
)

// FOR - neste ex.: enquanto esta cond for verdadeira, repita!
func main() {
	/*i := 0 //crie uma variável

	for i < 10 {
		i++
		fmt.Println("Incrementando i")
		time.Sleep(time.Second) //executa a variável, dorme um seg, executa a variável
	}

	fmt.Println(i)*/

	//OUTRA MANEIRA
	/*for j := 0; j < 10; j++ { //j += 2  se quiser incrementar de 2 em 2
		fmt.Println("Incrementando j", j)
		time.Sleep(time.Second)
	}*/

	//FOR COM RANGE
	nomes := [3]string{"João", "Davi", "Lucas"}

	for indice, nome := range nomes { //põe o índice e o valor
		fmt.Println(indice, nome)
	}

	/*for _, nome := range nomes { //põe o índice e o valor
		fmt.Println(nome)
	}
	SE QUISER QUE O INDICE NÃO APAREÇA _ */

	//ITERAR POR STRING
	for indice, letra := range "PALAVRA" {
		fmt.Println(indice, string(letra)) //se quiser printar a letra
	}
}
