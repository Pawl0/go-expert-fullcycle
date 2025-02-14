package main

import (
	"encoding/json"
	"os"
)

type Conta struct {
	Numero int `json:"-"`
	Saldo int `json:"saldo" validate:"gt=0"`
}

func main() {
	conta := Conta{Numero: 1, Saldo: 100}
	res, err := json.Marshal(conta)
	if err != nil {
		println(err)
	}
	println(string(res))

	err = json.NewEncoder(os.Stdout).Encode(conta)
	if err != nil {
		println(err)
	}

	jsonPuro:= []byte(`{"n": 2, "s": 200}`)
	var contaX Conta
	err = json.Unmarshal(jsonPuro, &contaX)
	if err != nil {
		println(err)
	}
	println(contaX.Numero, contaX.Saldo)
}