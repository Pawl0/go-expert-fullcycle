package main

import (
	"fmt"
	"go-expert-fullcycle/matematica"

	"github.com/google/uuid"
)

func main() {
	s := matematica.Soma(10,20)
	carro := matematica.Carro{Marca: "Ford"}

	fmt.Println(carro.Andar())
	fmt.Println("Resultado: ", s)
	fmt.Println(uuid.New())
}
