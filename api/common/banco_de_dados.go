package common

import (
	"github.com/yuriserka/Engenharia_de_Software/api/entidades"
)

var (
	// TabelaUsuario é uma stub para a tabela de Usuarios
	TabelaUsuario = make(map[string]*entidades.Usuario)
	// TabelaCartoesUsuario é uma stub para a relação de quais cartões de crédito um usuário possui
	TabelaCartoesUsuario = make(map[string]map[string]*entidades.CartaoDeCredito)
	// TabelaEventos é uma stub para a tabela de Eventos
	TabelaEventos = make(map[string]*entidades.Evento)
	// TabelaEventoUsuario é uma stub para a relação entre usuario e Evento
	TabelaEventoUsuario = make(map[string][]string)
	// TabelaApresentacoes é uma stub para a tabela de aresentações
	TabelaApresentacoes = make(map[string]*entidades.Apresentacao)
	// TabelaEventosApresentacoes é uma stub para a relação de evento com apresentação
	TabelaEventosApresentacoes = make(map[string]map[string]*entidades.Apresentacao)
)
