package controladoras

import (
	"errors"

	"github.com/yuriserka/Engenharia_de_Software/api/repositorios"

	"github.com/yuriserka/Engenharia_de_Software/api/utils"
)

// Autenticar verifica se os dados do usuário estão no banco de dados. Retorna se foi possível autenticar.
func Autenticar(cpf, senha string) error {
	if u := repositorios.GetUsuario(cpf); u != nil {
		if !utils.ValidaSenha(u.Senha, senha) {
			return errors.New("senha incorreta")
		}
		return nil
	}
	return errors.New("usuario nao cadastrado ainda")
}
