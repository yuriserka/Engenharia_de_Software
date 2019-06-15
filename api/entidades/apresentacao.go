package entidades

// Apresentacao é pelo o que os usuários irão pagar para ver.
type Apresentacao struct {
	Codigo, Sala    string
	Disponibilidade int
	Data, Horario   string
	Preco           float64
}
