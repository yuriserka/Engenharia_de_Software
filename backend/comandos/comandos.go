package comandos

import (
	"fmt"

	"github.com/Engenharia_de_Software/utils"

	ctrlServicos "github.com/Engenharia_de_Software/backend/controladoras_servicos"
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

// InterfaceApresentacao é uma Interface virtual, cujo as outras interfaces de apresentação terão como base
type InterfaceApresentacao interface {
	Executar(is *ctrlServicos.InterfaceServico)
}

// InterfaceApresentacaoCadastro é responsável por pegar o input do usuario e cadastra-lo no banco de dados
type InterfaceApresentacaoCadastro struct{}

// InterfaceApresentacaoLogin é responsável por pegar o input do usuario e verificar se o mesmo está cadastrado
// no banco de dados
type InterfaceApresentacaoLogin struct{}

// Executar (Cadastro) -> recebe o input do usuário, criptografa e é inserido no banco de dados pela Interface de Serviço.
func (ia *InterfaceApresentacaoCadastro) Executar(is ctrlServicos.InterfaceServico) {
	cpf, senha := getUserData()
	senhaCriptografada := utils.CriptografaSenha(string(senha))

	is.Executar(cpf, senhaCriptografada)
}

// Executar (Login) -> recebe o input do usuário, criptografa e é compara com o valor que está no banco de dados
// pela Interface de Serviço.
//
func (ia *InterfaceApresentacaoLogin) Executar(is ctrlServicos.InterfaceServico) {
	cpf, senha := getUserData()

	is.Executar(cpf, string(senha))
}
