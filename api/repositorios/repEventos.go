package repositorios

import (
	"errors"

	"github.com/yuriserka/Engenharia_de_Software/api/common"
	"github.com/yuriserka/Engenharia_de_Software/api/entidades"
)

// SetEvento insere um Evento no banco de dados
func SetEvento(codigo, nome, cidade, estado, tipo, classificacao string) {
	common.TabelaEventos[codigo] = &entidades.Evento{
		Codigo:        codigo,
		Nome:          nome,
		Cidade:        cidade,
		Estado:        estado,
		Tipo:          tipo,
		Classificacao: classificacao,
	}
}

// SetEventoUsuario tenta adicionar mais um evento para o usuario, se tiver mais que 5 retorna erro
func SetEventoUsuario(cpf, codigo string) error {
	if len(common.TabelaEventoUsuario[cpf]) > 5 {
		return errors.New("um usuário pode gerenciar somente 5 eventos por vez")
	}
	common.TabelaEventoUsuario[cpf] = append(common.TabelaEventoUsuario[cpf], codigo)
	return nil
}

// GetEvento retorna um Evento no banco de dados
func GetEvento(cod string) *entidades.Evento {
	return common.TabelaEventos[cod]
}

// GetTodosEventos retorna todos os Eventos do banco de dados
func GetTodosEventos() map[string]*entidades.Evento {
	return common.TabelaEventos
}

// UpdateEvento atualiza um Evento no banco de dados. Pode mudar somente o nome, classificação, estado e tipo
func UpdateEvento(cod, nome, class, estado, tipo string) {
	e := common.TabelaEventos[cod]
	e.Nome = nome
	e.Classificacao = class
	e.Estado = estado
	e.Tipo = tipo
	common.TabelaEventos[cod] = e
}
