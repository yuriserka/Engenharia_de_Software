package repositorios

import (
	"errors"

	"github.com/yuriserka/Engenharia_de_Software/api/common"
	"github.com/yuriserka/Engenharia_de_Software/api/entidades"
	"github.com/yuriserka/Engenharia_de_Software/api/utils"
)

// SetUsuario insere um usuario no banco de dados de usuarios
func SetUsuario(cpf, senha string) error {
	if len(cpf) == 0 || len(senha) == 0 {
		return errors.New("Cpf ou Senha com quantidade insuficiente de caracteres")
	}
	common.TabelaUsuario[cpf] = &entidades.Usuario{
		Cpf:   cpf,
		Senha: senha,
	}
	return nil
}

// GetUsuario retorna um usuario se ele existir no banco de dados de usuarios
func GetUsuario(cpf string) (*entidades.Usuario, error) {
	u, ok := common.TabelaUsuario[cpf]
	if !ok {
		return nil, errors.New("Não existe usuário cadastrado com este CPF")
	}
	return u, nil
}

// UpdateUsuario atualiza a senha do usuario
func UpdateUsuario(cpf, senha string) {
	u := common.TabelaUsuario[cpf]
	u.Senha = utils.CriptografaSenha([]byte(senha))
	common.TabelaUsuario[cpf] = u
}

// SetCartao insere um cartao de credito para o usuario
func SetCartao(cpf, num, cod, validade string) error {
	mapCartao, existe := common.TabelaCartoesUsuario[cpf]
	if !existe {
		mapCartao = make(map[string]*entidades.CartaoDeCredito)
		common.TabelaCartoesUsuario[cpf] = mapCartao
	}
	if _, ok := mapCartao[num]; !ok {
		mapCartao[num] = &entidades.CartaoDeCredito{
			Numero:   num,
			Codigo:   cod,
			Validade: validade,
		}
	} else {
		return errors.New("Cartão ja cadastrado")
	}
	return nil
}

// GetCartao retorna um dos cartoes de credito do usuario
func GetCartao(cpf, num string) (*entidades.CartaoDeCredito, error) {
	cartoes, ok := common.TabelaCartoesUsuario[cpf]
	if !ok {
		return nil, errors.New("Cartão de Crédito não Cadastrado")
	}
	return cartoes[num], nil
}

// UpdateCartao atualiza a data de validade do cartao
func UpdateCartao(cpf, num, validade string) {
	cartoes, _ := common.TabelaCartoesUsuario[cpf]
	c, _ := cartoes[num]
	c.Validade = validade
	cartoes[num] = c
}

// GetCartoesUsuario retorna todos os cartoes de credito do usuario
func GetCartoesUsuario(cpf string) (map[string]*entidades.CartaoDeCredito, error) {
	v, _ := common.TabelaCartoesUsuario[cpf]
	if len(v) == 0 {
		return nil, errors.New("Este Usuário não possui Cartões de Crédito cadastrados")
	}
	return v, nil
}

// GetEventosUsuario retorna todos os codigos de eventos criados pelo usuario
func GetEventosUsuario(cpf string) ([]string, error) {
	es, ok := common.TabelaEventoUsuario[cpf]
	if !ok {
		return nil, errors.New("Este usuário não está gerenciando nenhum Evento")
	}
	return es, nil
}
