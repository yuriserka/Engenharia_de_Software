package servicos

import (
	"fmt"
	"time"

	"github.com/yuriserka/Engenharia_de_Software/entidades"
)

var banco = make(map[string]string)

// ControladoraServicoCadastro é responsável por cadastrar um novo usuário ao sistema.
//
// Implementa a interfaceServicoCadastro.
type ControladoraServicoCadastro struct{}

// CadastrarNovoUsuario cadastra um novo usuário no sistema se ele não existir.
func (ctrlSCadastro *ControladoraServicoCadastro) CadastrarNovoUsuario(cpf, senha string) bool {
	if _, exist := banco[cpf]; exist {
		fmt.Println("Usuario ja cadastrado no sistema.")
		return false
	}

	banco[cpf] = senha
	fmt.Printf("inserindo usuario: %#v no banco de dados\n", entidades.Usuario{Cpf: cpf, Senha: senha})
	time.Sleep(2 * time.Second)

	return true
}
