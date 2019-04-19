package servicos

import (
	"fmt"

	"github.com/Engenharia_de_Software/utils"
)

// ControladoraServicoLogin é responsável pela autenticação de usuários no sistema.
//
// Implementa a InterfaceServicoLogin
type ControladoraServicoLogin struct{}

// Autenticar verifica se os dados do usuário estão no banco de dados. Retorna o cpf digitado e se foi possível
// autenticar.
func (ctrlSLogin *ControladoraServicoLogin) Autenticar(cpf, senha string) bool {
	if _, exist := banco[cpf]; !exist {
		fmt.Println("Usuario nao cadastrado ainda.")
		return false
	}

	if s, _ := banco[cpf]; !utils.ValidaSenha(s, senha) {
		fmt.Println("Senha incorreta.")
		return false
	}

	return true
}
