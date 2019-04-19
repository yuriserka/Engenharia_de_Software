package servicos

import (
	"fmt"
	"time"

	"github.com/Engenharia_de_Software/utils"
)

var banco = make(map[string]string)

// InterfaceServico é a interface que se resopnsabiliza pela execução da lógica de negócio do sistema
//
// Executar é o método que irá ser responsável por realizar determinada operação do backend,
// por esse motivo é um variadico de interface{}, pois todas as interfaces de serviço tomaram esta como base.
type InterfaceServico interface {
	Executar(args ...interface{}) bool
}

// InterfaceServicoCadastro é a interface responsável pelo cadastro de um novo usuário ao sistema.
type InterfaceServicoCadastro struct{}

// InterfaceServicoLogin é a interface responsável pelo login dos usuários do sistema.
type InterfaceServicoLogin struct{}

// Executar recebe como argumento (cpf, senha string) e
// retorna true se foi possivel cadastrar ou false caso contrário.
func (isk *InterfaceServicoCadastro) Executar(args ...interface{}) bool {
	// conectar no banco de dados e blabla
	if len(args) != 2 {
		fmt.Println("é necessario do cpf e senha apenas.")
		return false
	}

	// https://tour.golang.org/methods/15
	cpf, senha := args[0].(string), args[1].(string)
	banco[cpf] = senha

	time.Sleep(1 * time.Second)

	return true
}

// Executar recebe como argumento (cpf, senha string) e
// retorna true se as senhas forem iguais ou false caso contrário.
func (isl *InterfaceServicoLogin) Executar(args ...interface{}) bool {
	if len(args) != 2 {
		fmt.Println("é necessario do cpf e senha apenas.")
		return false
	}

	cpf, senha := args[0].(string), args[1].(string)

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
