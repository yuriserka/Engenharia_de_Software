package controladoras

import (
	"fmt"

	"github.com/yuriserka/Engenharia_de_Software/api/repositorios"
)

// CadastrarNovoUsuario cadastra um novo usuário no sistema se ele não existir.
func CadastrarNovoUsuario(cpf, senha string) bool {
	if u := repositorios.GetUsuario(cpf); u == nil {
		repositorios.SetUsuario(cpf, senha)
		return true
	}
	fmt.Println("ja existe usuario cadastrado com este CPF")
	return false
}
