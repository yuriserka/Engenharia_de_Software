package controladoras

import (
	"fmt"

	"github.com/yuriserka/Engenharia_de_Software/api/repositorios"
	"github.com/yuriserka/Engenharia_de_Software/api/utils"
)

// CadastrarNovoUsuario cadastra um novo usuário no sistema se ele não existir.
func CadastrarNovoUsuario(cpf string, senha []byte) bool {
	senhaCriptografada := utils.CriptografaSenha(senha)
	if u := repositorios.GetUsuario(cpf); u == nil {
		repositorios.SetUsuario(cpf, string(senhaCriptografada))
		return true
	}
	fmt.Println("ja existe usuario cadastrado com este CPF")
	return false
}
