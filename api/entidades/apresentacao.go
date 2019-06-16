package entidades

import "fmt"

// Apresentacao é pelo o que os usuários irão pagar para ver.
type Apresentacao struct {
	Codigo, Sala    string
	Disponibilidade int
	Data, Horario   string
	Preco           float64
}

// ToString facilita a impressao desta entidade
func (apr Apresentacao) ToString() string {
	return fmt.Sprintf("Codigo: %s\nSala: %s\nPreco: %f\nData: %sh\nDisponibilidade: %d", apr.Codigo,
		apr.Sala, apr.Preco, apr.Data+" as "+apr.Horario, apr.Disponibilidade)
}
