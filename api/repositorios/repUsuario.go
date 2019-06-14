package repositorios

import (
	"time"

	"github.com/yuriserka/Engenharia_de_Software/api/common"
	"github.com/yuriserka/Engenharia_de_Software/api/entidades"
)

// SetUsuario insere um usuario no banco de dados de usuarios
func SetUsuario(cpf, senha string) {
	common.TabelaUsuario[cpf] = &entidades.Usuario{Cpf: cpf, Senha: senha}
	time.Sleep(1 * time.Second)
}

// GetUsuario retorna um usuario se ele existir no banco de dados de usuarios
func GetUsuario(cpf string) *entidades.Usuario {
	u, _ := common.TabelaUsuario[cpf]
	return u
}
