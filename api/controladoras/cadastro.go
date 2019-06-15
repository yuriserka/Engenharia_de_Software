package controladoras

import (
	"errors"

	"github.com/yuriserka/Engenharia_de_Software/api/repositorios"
	"github.com/yuriserka/Engenharia_de_Software/api/utils"
)

// CadastrarNovoUsuario cadastra um novo usuário no sistema se ele não existir.
func CadastrarNovoUsuario(cpf string, senha []byte) error {
	senhaCriptografada := utils.CriptografaSenha(senha)
	if u := repositorios.GetUsuario(cpf); u == nil {
		repositorios.SetUsuario(cpf, string(senhaCriptografada))
		return nil
	}

	return errors.New("ja existe usuario cadastrado com este CPF")
}
