package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println("Arrays e slices")

	//ARRAY = LISTA DE VALORES
	var array1 [5]string    //tds os dados do array do msm tipo
	array1[0] = "Posição 1" //colocar um valor na primeira posição
	fmt.Println(array1)

	//outra maneira
	array2 := [5]string{"Posição 1", "Posição 2", "Posição 3", "Posição 4", "Posição 5"}
	fmt.Println(array2)

	//SLICE ñ estabelece tamanho, pode mudar de acordo com a necessidade, mas só pode ter dados do msm tipo
	slice := []int{10, 12, 13, 14, 15, 16, 17}
	fmt.Println(slice)

	//devolve o tipo de uma variável
	fmt.Println(reflect.TypeOf(slice))
	fmt.Println(reflect.TypeOf(array2))

	//append = coloca um item novo no slice e retorna um novo com o item adicionado
	//no ex abaixo só estou alterando/adc valor o slice
	slice = append(slice, 18)
	fmt.Println(slice)

	//esse slice é um pedaço do array2 e dentro do colchete põe os intervalos do índices que deseja
	slice2 := array2[1:3]
	fmt.Println(slice2)

	//ARRAYS INTERNOS
	slice3 := make([]float32, 10, 11) //tipo, tamanho(qnts itens tem) e capacidade(qnts consegue ter)
	fmt.Println(slice3)
	//saber tamanho e capacidade
	fmt.Println(len(slice3)) //length tamanho
	fmt.Println(cap(slice3)) //cap capacidade

	slice3 = append(slice3, 5)
	fmt.Println(slice3)

	//quando ultrapassa a capacidade, ele dobra de tamanho e gera um novo slice
	slice4 := make([]float32, 5)
	fmt.Println(slice4)
	slice4 = append(slice4, 10)
	fmt.Println(len(slice4))
	fmt.Println(cap(slice4))

}
