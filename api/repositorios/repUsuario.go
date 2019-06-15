package repositorios

import (
	"github.com/yuriserka/Engenharia_de_Software/api/common"
	"github.com/yuriserka/Engenharia_de_Software/api/entidades"
	"github.com/yuriserka/Engenharia_de_Software/api/utils"
)

// SetUsuario insere um usuario no banco de dados de usuarios
func SetUsuario(cpf, senha string) {
	common.TabelaUsuario[cpf] = &entidades.Usuario{Cpf: cpf, Senha: senha}
}

// GetUsuario retorna um usuario se ele existir no banco de dados de usuarios
func GetUsuario(cpf string) *entidades.Usuario {
	u, _ := common.TabelaUsuario[cpf]
	return u
}

// UpdateUsuario atualiza a senha do usuario
func UpdateUsuario(cpf, senha string) {
	u := common.TabelaUsuario[cpf]
	u.Senha = utils.CriptografaSenha([]byte(senha))
	common.TabelaUsuario[cpf] = u
}

// SetCartao insere um cartao de credito para o usuario
func SetCartao(cpf, num, cod, validade string) {
	common.TabelaCartoesUsuario[cpf][num] = &entidades.CartaoDeCredito{
		Numero:   num,
		Codigo:   cod,
		Validade: validade,
	}
}

// GetCartao retorna um dos cartoes de credito do usuario
func GetCartao(cpf, num string) *entidades.CartaoDeCredito {
	cartoes, _ := common.TabelaCartoesUsuario[cpf]
	c, _ := cartoes[num]
	return c
}

// UpdateCartao atualiza a data de validade do cartao
func UpdateCartao(cpf, num, validade string) {
	cartoes, _ := common.TabelaCartoesUsuario[cpf]
	c, _ := cartoes[num]
	c.Validade = validade
	cartoes[num] = c
}

// GetCartoesUsuario retorna todos os cartoes de credito do usuario
func GetCartoesUsuario(cpf string) map[string]*entidades.CartaoDeCredito {
	cs, _ := common.TabelaCartoesUsuario[cpf]
	return cs
}
