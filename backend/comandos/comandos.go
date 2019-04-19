package comandos

import (
	"fmt"

	"github.com/Engenharia_de_Software/utils"

	servicos "github.com/Engenharia_de_Software/backend/controladoras_servicos"
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

// InterfaceApresentacao é a interface que se resopnsabiliza pela execução da parte textual, ou telas, do programa
//
// Executar é o método que irá ser responsável por realizar determinada operação do frontend,
// por esse motivo é preciso de uma interface de serviço, pois será ela que irá concretizar a ação.
// Todas as interfaces de apresentação tomaram esta como base.
type InterfaceApresentacao interface {
	Executar(is servicos.InterfaceServico) bool
}

// InterfaceApresentacaoCadastro é a interface responsável pela solicitação dos dados do novo usuário do sistema.
type InterfaceApresentacaoCadastro struct{}

// InterfaceApresentacaoLogin é a interface responsável pela solicitação dos dados de um usuário para fazer login no sistema.
type InterfaceApresentacaoLogin struct{}

// Executar tenta cadastrar determinado usuário ao sistema.
func (iak *InterfaceApresentacaoCadastro) Executar(is servicos.InterfaceServico) bool {
	cpf, senha := getUserData()
	senhaCriptografada := utils.CriptografaSenha(string(senha))

	return is.Executar(cpf, senhaCriptografada)
}

// Executar tenta entrar no sistema dado o CPF e a senha de um usuário.
func (ial *InterfaceApresentacaoLogin) Executar(is servicos.InterfaceServico) bool {
	cpf, senha := getUserData()

	return is.Executar(cpf, string(senha))
}
