package entidades

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
