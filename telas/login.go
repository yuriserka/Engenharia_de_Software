package telas

import (
	"fmt"

	"github.com/yuriserka/Engenharia_de_Software/interfaces"
	"github.com/yuriserka/Engenharia_de_Software/utils"
	"github.com/howeyc/gopass"
)

var getUserData = func() (cpf string, senha []byte) {
	fmt.Printf("Cpf: ")
	fmt.Scanf("%s\n", &cpf)

	fmt.Printf("Senha: ")
	senha, err := gopass.GetPasswdMasked()

	if err != nil {
		panic(err.Error())
	}

	return
}

// ControladoraApresentacaoLogin é responsável por apresentar a tela de login do sistema.
//
// Implementa os métodos da InterfaceServicoLogin
type ControladoraApresentacaoLogin struct {
	ctrlServicoLogin interfaces.InterfaceServicoLogin
}

// SetControladoraServicoLogin mantém uma instância da interface de serviço, para que a controladora
// possa se comunicar com os serviços providos pela camada de serviços.
func (ctrlALogin *ControladoraApresentacaoLogin) SetControladoraServicoLogin(ISLogin interfaces.InterfaceServicoLogin) {
	ctrlALogin.ctrlServicoLogin = ISLogin
}

// Autenticar solicita os dados do usuário e verifica se o mesmo está no bando de dados. Se estiver então é retornado
// o cpf dele e True, caso contrário retorna uma string vazia e False.
func (ctrlALogin *ControladoraApresentacaoLogin) Autenticar() (string, bool) {
	utils.ClearScreen()
	fmt.Println("\tEntrar")

	cpf, senha := getUserData()

	if logado := ctrlALogin.ctrlServicoLogin.Autenticar(cpf, string(senha)); !logado {
		utils.Pause()
		return "", false
	}

	return cpf, true
}
