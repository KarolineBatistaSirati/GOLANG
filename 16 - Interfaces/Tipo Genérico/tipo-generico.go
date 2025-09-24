package main

import "fmt"

func generica(interf interface{}) {
	fmt.Println(interf)

}

func main() {
	generica("String")
	generica(1)
	generica(true)

	//O PRINTLN FUNCIONA ASSIM, RECEBE QLQR TIPO DE DADO
	fmt.Println(1, 2, "String", false, float64(12344))

	//IGNORANDO O TIPO, MAS NÃO É RECOMENDADO!!!!
	/*mapa := map[interface{}]interface{}{
	1: "String",
	float32(100): true,
	"String": "String",

	}

	fmt.Println(mapa)*/
}
