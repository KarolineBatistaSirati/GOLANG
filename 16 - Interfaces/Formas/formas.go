package main

import (
	"fmt"
	"math"
)

type retangulo struct {
	altura  float64
	largura float64
}

func (r retangulo) area() float64 {
	return r.altura * r.largura
}

type circulo struct {
	raio float64
}

// COMO CRIAR INTERFACE (parecida com struct, mas ao invés de colocar struct, põe interface
type forma interface {
	area() float64
}

// dps cria uma funcao
func escreverArea(f forma) {
	fmt.Printf("A área da forma é %0.2f\n", f.area()) //USE PRINTF
}

func (c circulo) area() float64 {
	return math.Pi * (c.raio * c.raio)
}

// USADAS QUANDO NECESSÁRIO TER FLEXIBILIDADE COM TIPOS
func main() {
	r := retangulo{10, 15} //criei uma variavel e tribuí valor
	escreverArea(r)        //chamar a funcao passando o retangulo

	c := circulo{10}
	escreverArea(c)
}
