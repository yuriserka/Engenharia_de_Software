package entidades

import "time"

// Usuario é o principal ator do sistema.
//
// Qualquer usuário pode obter dados sobre os eventos, adicionar ou editar um evento que ele criou e
// comprar ingressos para outros.
//
// Ao ser cadastrado, pode cadastrar, descadastrar e alterar evento, além de poder retirar sua conta do sistema.
//
// Só é permitido a gerencia de 5 eventos.
type Usuario struct {
	Cpf, Senha string
}

// Evento é onde ocorrerá as apresentações
//
// Pode ter os seguintes tipos: teatro, esporte, show nacional e internacional.
//
// Não pode ser descadastrado se já tiver sido vendido um ingresso para alguma de suas apresentações.
//
// Só é permitida a existência de 10 apresentações
type Evento struct {
	Codigo, Nome, Cidade, Estado, Tipo, Classificaçao string
}

// Apresentacao é pelo o que os usuários irão pagar para ver.
type Apresentacao struct {
	Codigo, Sala, Disponibilidade string
	Data, Horario                 time.Time
	Preco                         float64
}

// Ingresso é o meio que os usuários tem para poder gerenciar as vendas de cada evento.
type Ingresso struct {
	Codigo string
}

// CartaoDeCredito é a classe que servirá para validar se o usuário pode ou não pagar pela apresentação.
type CartaoDeCredito struct {
	Numero, Codigo, Validade string
}
