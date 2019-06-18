package repositorios

import (
	"errors"

	"github.com/yuriserka/Engenharia_de_Software/api/common"
	"github.com/yuriserka/Engenharia_de_Software/api/entidades"
)

// SetEvento insere um Evento no banco de dados
func SetEvento(codigo, nome, cidade, estado, tipo, classificacao string) error {
	if _, ok := common.TabelaEventos[codigo]; ok {
		return errors.New("Ja existe um evento com este código")
	}
	common.TabelaEventos[codigo] = &entidades.Evento{
		Codigo:        codigo,
		Nome:          nome,
		Cidade:        cidade,
		Estado:        estado,
		Tipo:          tipo,
		Classificacao: classificacao,
	}
	return nil
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
func GetEvento(cod string) (*entidades.Evento, error) {
	v, ok := common.TabelaEventos[cod]
	if !ok {
		return nil, errors.New("Evento não cadastrado")
	}
	return v, nil
}

// GetTodosEventos retorna todos os Eventos do banco de dados
func GetTodosEventos() (map[string]*entidades.Evento, error) {
	v := common.TabelaEventos
	if len(v) == 0 {
		return nil, errors.New("Não há nenhum evento cadastrado ainda")
	}
	return v, nil
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
func GetApresentacao(cod string) (*entidades.Apresentacao, error) {
	v, ok := common.TabelaApresentacoes[cod]
	if !ok {
		return nil, errors.New("Apresentação não cadastrada")
	}
	return v, nil
}

// SetApresentacaoEvento faz a relação entre envento e apresentação
func SetApresentacaoEvento(codigoEvento string, a *entidades.Apresentacao) error {
	mapApresentacao, existe := common.TabelaEventosApresentacoes[codigoEvento]
	if !existe {
		mapApresentacao = make(map[string]*entidades.Apresentacao)
		common.TabelaEventosApresentacoes[codigoEvento] = mapApresentacao
	}
	if len(mapApresentacao) > 10 {
		return errors.New("Um evento deve possuir apenas 10 apresentações")
	}
	if _, ok := mapApresentacao[a.Codigo]; !ok {
		mapApresentacao[a.Codigo] = a
	} else {
		return errors.New("Apresentação ja cadastrada")
	}
	return nil
}

// GetApresentacoes retorna todas as apresentações do banco de dados
func GetApresentacoes(codigoEvento string) (map[string]*entidades.Apresentacao, error) {
	if _, e := GetEvento(codigoEvento); e != nil {
		return nil, e
	}
	v, ok := common.TabelaEventosApresentacoes[codigoEvento]
	if !ok {
		return nil, errors.New("Evento não possui nenhuma apresentação ainda")
	}
	return v, nil
}
