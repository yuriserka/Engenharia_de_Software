package controladoras

import (
	"fmt"

	"github.com/yuriserka/Engenharia_de_Software/api/repositorios"

	"github.com/yuriserka/Engenharia_de_Software/api/utils"
)

// Autenticar verifica se os dados do usuário estão no banco de dados. Retorna o cpf digitado e se foi possível
// autenticar.
func Autenticar(cpf, senha string) bool {
	if u := repositorios.GetUsuario(cpf); u != nil {
		if !utils.ValidaSenha(u.Senha, senha) {
			fmt.Println("Senha incorreta.")
			return false
		}
		return true
	}
	fmt.Println("Usuario nao cadastrado ainda.")
	return false
}
