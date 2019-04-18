package ctrlServicos

import (
	"fmt"
	"time"

	"github.com/Engenharia_de_Software/utils"
)

type InterfaceServico interface {
	// Cada interface terá seu método de executar, por esse motivo é variadico
	Executar(args ...interface{}) bool
}

var banco = make(map[string]string)

type InterfaceServicoCadastro struct{}
type InterfaceServicoLogin struct{}

// Executar registra um usuário no banco de dados.
// args = (cpf, senha)
func (is *InterfaceServicoCadastro) Executar(args ...interface{}) bool {
	// conectar no banco de dados e blabla
	if len(args) != 2 {
		fmt.Println("é necessario do cpf e senha apenas.")
		return false
	}

	// https://tour.golang.org/methods/15
	cpf, senha := args[0].(string), args[1].(string)

	banco[cpf] = senha

	time.Sleep(2 * time.Second)
	fmt.Println("Tudo certo, usuario cadastrado")

	for key, val := range banco {
		fmt.Printf("CPF: %s, senha: %s\n", key, val)
	}
	return true
}

// Executar verifica se existe um usuário com o cpf no banco de dados, se houver então
// é retornado sua senha e verificado com a passada como argumento, se forem iguais, então
// o usuário entra no sistema, caso contrário é informado "Senha incorreta".
//
// Caso não haja o cpf
// passado no banco de dados, então informa "Usuario nao cadastrado ainda.".
func (is *InterfaceServicoLogin) Executar(args ...interface{}) bool {
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
