package controladoras

import (
	"github.com/yuriserka/Engenharia_de_Software/api/entidades"
	"github.com/yuriserka/Engenharia_de_Software/api/repositorios"
)

// CadastrarNovoEvento cadastra novos eventos e os associa ao usuario criador
func CadastrarNovoEvento(cpf, codigo, nome, cidade, estado, tipo, classificacao string) error {
	if err := repositorios.SetEvento(codigo, nome, cidade, estado, tipo, classificacao); err != nil {
		return err
	}
	return repositorios.SetEventoUsuario(cpf, codigo)
}

// RecuperarEventos Mostra informação sobre todos os eventos cadastrados
func RecuperarEventos() (map[string]*entidades.Evento, error) {
	return repositorios.GetTodosEventos()
}

// RecuperarEvento retorna um dado evento
func RecuperarEvento(cod string) (*entidades.Evento, error) {
	return repositorios.GetEvento(cod)
}

// CadastrarNovaApresentacao insere uma nova apresentação no banco de dados
func CadastrarNovaApresentacao(codigo, sala string, disponibilidade int, data, horario string, preco float64) error {
	return repositorios.SetApresentacao(codigo, sala, disponibilidade, data, horario, preco)
}

// LigarApresentacaoEvento relaciona o evento com a apresentação
func LigarApresentacaoEvento(codEvento string, apr *entidades.Apresentacao) error {
	return repositorios.SetApresentacaoEvento(codEvento, apr)
}

// RecuperarApresentacoes retorna todas as apresentações de um evento
func RecuperarApresentacoes(codEvento string) (map[string]*entidades.Apresentacao, error) {
	return repositorios.GetApresentacoes(codEvento)
}

// RecuperarApresentacao Retorna apenas uma apresentação
func RecuperarApresentacao(cod string) (*entidades.Apresentacao, error) {
	return repositorios.GetApresentacao(cod)
}
