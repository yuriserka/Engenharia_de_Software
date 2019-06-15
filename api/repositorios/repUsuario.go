package repositorios

import (
	"github.com/yuriserka/Engenharia_de_Software/api/common"
	"github.com/yuriserka/Engenharia_de_Software/api/entidades"
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

// SetCartao insere um cartao de credito para o usuario
func SetCartao(cpf, num, cod, validade string) {
	common.TabelaCartoesUsuario[cpf] = append(common.TabelaCartoesUsuario[cpf],
		&entidades.CartaoDeCredito{
			Numero:   num,
			Codigo:   cod,
			Validade: validade,
		})
}

// GetCartao retorna um dos cartoes de credito do usuario
func GetCartao(cpf, num string) *entidades.CartaoDeCredito {
	arrCartao, _ := common.TabelaCartoesUsuario[cpf]
	for _, c := range arrCartao {
		if c.Numero == num {
			return c
		}
	}
	return nil
}

// GetCartoesUsuario retorna todos os cartoes de credito do usuario
func GetCartoesUsuario(cpf string) []*entidades.CartaoDeCredito {
	cs, _ := common.TabelaCartoesUsuario[cpf]
	return cs
}
