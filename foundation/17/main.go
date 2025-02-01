package main

import "fmt"

type Client struct {
	nome string
}

func (c Client) andou() {
	c.nome = "Wesley Willians"
	fmt.Printf("O cliente %v andou\n", c.nome)
}

type Conta struct {
	saldo int
}

func NewConta() *Conta {
	return &Conta{saldo: 0}
}

func (c *Conta) simular(valor int) int {
	c.saldo += valor
	println(c.saldo)
	return c.saldo
}

func main() {
	conta := Conta{saldo: 100}
	conta.simular((200))
	println(conta.saldo)
}
