package matematica

func Soma[T int | float64](a, b T) T {
	return a + b
}

type Carro struct{
	Marca string
}

func (c Carro) Andar() string {
	return "Carro andando"
}