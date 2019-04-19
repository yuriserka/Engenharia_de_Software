package telas

import (
	"fmt"

	"github.com/Engenharia_de_Software/interfaces"
	"github.com/Engenharia_de_Software/utils"
)

// ControladoraApresentacaoCadastro é responsável por apresentar a tela de cadastro no sistema.
//
// Implementa os métodos da InterfaceServicoCadastro
type ControladoraApresentacaoCadastro struct {
	ctrlServicoCadastro interfaces.InterfaceServicoCadastro
}

// SetControladoraServicoCadastro mantém uma instância da interface de serviço, para que a controladora
// possa se comunicar com os serviços providos pela camada de serviços.
func (ctrlACadastro *ControladoraApresentacaoCadastro) SetControladoraServicoCadastro(
	ISCadastro interfaces.InterfaceServicoCadastro) {
	ctrlACadastro.ctrlServicoCadastro = ISCadastro
}

// Cadastrar é responsável por pegar os dados do usuário e mandar a controladora de serviços inserir
// no banco de dados
func (ctrlACadastro *ControladoraApresentacaoCadastro) Cadastrar() {
	utils.ClearScreen()
	fmt.Println("\tCadastre-se")

	cpf, senha := getUserData()
	senhaCriptografada := utils.CriptografaSenha(senha)

	ctrlACadastro.ctrlServicoCadastro.CadastrarNovoUsuario(cpf, senhaCriptografada)
}
