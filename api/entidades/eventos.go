package entidades

import "fmt"

// Evento é onde ocorrerá as apresentações
//
// Pode ter os seguintes tipos: teatro, esporte, show nacional e internacional.
//
// Não pode ser descadastrado se já tiver sido vendido um ingresso para alguma de suas apresentações.
//
// Só é permitida a existência de 10 apresentações
type Evento struct {
	Codigo, Nome, Cidade, Estado, Tipo, Classificacao string
}

// ToString facilita a impressão desta entidade
func (e Evento) ToString() string {
	return fmt.Sprintf("Codigo: %s\nNome: %s\nLocal: %s\nClassificação: %s\nTipo: %s", e.Codigo,
		e.Nome, e.Cidade+"-"+e.Estado, e.Classificacao, e.Tipo)
}
