package telas

import (
	"fmt"

	"github.com/yuriserka/Engenharia_de_Software/interfaces"
	"github.com/yuriserka/Engenharia_de_Software/utils"
)

// ControladoraApresentacaoControle responsável por controlar o fluxo do programa de forma geral, ela que irá controlar os
// dois principais menus que o sistema possui.
//
// Implementa os métodos da InterfaceServicoControle
type ControladoraApresentacaoControle struct {
	iALogin    interfaces.InterfaceApresentacaoLogin
	iACadastro interfaces.InterfaceApresentacaoCadastro
	iAUsuario  interfaces.InterfaceApresentacaoUsuario
	iAEventos  interfaces.InterfaceApresentacaoEventos
}

// Init apresentará as opções de Login e Cadastro. É onde o programa começa.
func (ctrlAControle *ControladoraApresentacaoControle) Init() {
	var opt int
	for opt != 3 {
		utils.ClearScreen()
		fmt.Print("\t\tSistema de Vendas de Ingressos\n", "Escolha uma das opções abaixo:\n")
		fmt.Print("[1] login\n", "[2] cadastrar-se\n", "[3] sair\n", "\topção: ")

		switch fmt.Scanf("%d\n", &opt); opt {
		case 1:
			if cpf, logado := ctrlAControle.iALogin.Autenticar(); logado {
				ctrlAControle.controleLogado(cpf)
			}
		case 2:
			ctrlAControle.iACadastro.Cadastrar()
			utils.Pause()
		}
	}
}

// ControleLogado apresentará as opções de gestão de Usuário e de Eventos. Só é acessível após o usuário ser
// autenticado.
func (ctrlAControle *ControladoraApresentacaoControle) controleLogado(cpf string) {
	var opt int
	for opt != 3 {
		utils.ClearScreen()
		fmt.Print("\tBem vindo.\n", "Escolha uma das opções abaixo:\n")
		fmt.Print("[1] Informacoes do Usuario\n", "[2] Visualizar Eventos\n", "[3] voltar\n", "\topção: ")

		switch fmt.Scanf("%d\n", &opt); opt {
		case 1:
			utils.Pause()
		case 2:
			utils.Pause()
		}
	}
}

// SetControladoraApresentacaoLogin recebe uma controladora de apresentação do Login.
func (ctrlAControle *ControladoraApresentacaoControle) SetControladoraApresentacaoLogin(
	IALogin interfaces.InterfaceApresentacaoLogin) {
	ctrlAControle.iALogin = IALogin
}

// SetControladoraApresentacaoCadastro recebe uma controladora de apresentação de Cadastro.
func (ctrlAControle *ControladoraApresentacaoControle) SetControladoraApresentacaoCadastro(
	IACadastro interfaces.InterfaceApresentacaoCadastro) {
	ctrlAControle.iACadastro = IACadastro
}

// SetControladoraApresentacaoUsuario recebe uma controladora de apresentação de Gestão de Usuário.
func (ctrlAControle *ControladoraApresentacaoControle) SetControladoraApresentacaoUsuario(
	IAUsuario interfaces.InterfaceApresentacaoUsuario) {
	ctrlAControle.iAUsuario = IAUsuario
}

// SetControladoraApresentacaoEventos recebe uma controladora de apresentação de Gestão de Eventos.
func (ctrlAControle *ControladoraApresentacaoControle) SetControladoraApresentacaoEventos(
	IAEventos interfaces.InterfaceApresentacaoEventos) {
	ctrlAControle.iAEventos = IAEventos
}
