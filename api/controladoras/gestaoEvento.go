package controladoras

import (
	"github.com/yuriserka/Engenharia_de_Software/api/entidades"
	"github.com/yuriserka/Engenharia_de_Software/api/repositorios"
)

// CadastrarNovoEvento cadastra novos eventos e os associa ao usuario criador
func CadastrarNovoEvento(cpf, codigo, nome, cidade, estado, tipo, classificacao string) error {
	repositorios.SetEvento(codigo, nome, cidade, estado, tipo, classificacao)
	err := repositorios.SetEventoUsuario(cpf, codigo)

	return err
}

// RecuperarEventos Mostra informação sobre todos os eventos cadastrados
func RecuperarEventos() map[string]*entidades.Evento {
	return repositorios.GetTodosEventos()
}
