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
	common.TabelaEventoUsuario[cpf] = append(common.TabelaEventoUsuario[cpf], codigo)
	size := len(common.TabelaEventoUsuario[cpf])
	if size > 5 {
		common.TabelaEventoUsuario[cpf] = common.TabelaEventoUsuario[cpf][:size-1]
		return errors.New("um usuário pode gerenciar somente 5 eventos por vez")
	}
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

// SetApresentacao insere uma apresentação no banco de dados
func SetApresentacao(codigo, sala string, disponibilidade int, data, horario string, preco float64) error {
	if a, _ := common.TabelaApresentacoes[codigo]; a != nil {
		return errors.New("Ja existe uma apresentação com este codigo")
	}
	common.TabelaApresentacoes[codigo] = &entidades.Apresentacao{
		Codigo:          codigo,
		Sala:            sala,
		Disponibilidade: disponibilidade,
		Data:            data,
		Horario:         horario,
		Preco:           preco,
	}
	return nil
}

// GetApresentacao retorna uma apresentação do banco de dados
func GetApresentacao(cod string) *entidades.Apresentacao {
	return common.TabelaApresentacoes[cod]
}

// SetApresentacaoEvento faz a relação entre envento e apresentação
func SetApresentacaoEvento(codigoEvento string, a *entidades.Apresentacao) error {
	mapApresentacao, existe := common.TabelaEventosApresentacoes[codigoEvento]
	if !existe {
		mapApresentacao = make(map[string]*entidades.Apresentacao)
		common.TabelaEventosApresentacoes[codigoEvento] = mapApresentacao
	}
	if _, ok := mapApresentacao[a.Codigo]; !ok {
		mapApresentacao[a.Codigo] = a
	} else {
		return errors.New("Apresentação ja cadastrada")
	}
	return nil
}

// GetApresentacoes retorna todas as apresentações do banco de dados
func GetApresentacoes(codigoEvento string) map[string]*entidades.Apresentacao {
	return common.TabelaEventosApresentacoes[codigoEvento]
}
