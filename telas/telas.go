package telas

import (
	"fmt"

	"github.com/Engenharia_de_Software/backend/comandos"
	servicos "github.com/Engenharia_de_Software/backend/controladoras_servicos"
	"github.com/Engenharia_de_Software/utils"
)

// TelaPrincipal apresentará as opções de Login e Cadastro. Cada uma dessas opções levará a outra tela,
// estas que comporão a aplicação.
func TelaPrincipal() {
	var opt int
	for opt != 3 {
		utils.ClearScreen()
		fmt.Print("\t\tSistema de Vendas de Ingressos\n", "Escolha uma das opções abaixo:\n")
		fmt.Print("[1] login\n", "[2] cadastrar-se\n", "[3] sair\n", "\topção: ")

		switch fmt.Scanf("%d\n", &opt); opt {
		case 1:
			if logado := telaLogin(); logado {
				telaLogado()
			} else {
				utils.Pause()
			}
		case 2:
			telaCadastro()
			utils.Pause()
		}
	}
}

func telaLogin() bool {
	utils.ClearScreen()
	fmt.Println("\tEntrar")

	loginInterface := new(comandos.InterfaceApresentacaoLogin)
	servico := &servicos.InterfaceServicoLogin{}

	return loginInterface.Executar(servico)
}

func telaCadastro() {
	utils.ClearScreen()
	fmt.Println("\tCadastre-se")

	registerInterface := new(comandos.InterfaceApresentacaoCadastro)
	servico := &servicos.InterfaceServicoCadastro{}

	if cadastradoComSucesso := registerInterface.Executar(servico); cadastradoComSucesso {
		fmt.Println("Cadastrado com sucesso")
	}
}

func telaLogado() {
	var opt int
	for opt != 3 {
		utils.ClearScreen()
		fmt.Print("\t\tSistema de Vendas de Ingressos\n", "Escolha uma das opções abaixo:\n")
		fmt.Print("[1] Informacoes do Usuario\n", "[2] Visualizar Eventos\n", "[3] voltar\n", "\topção: ")

		switch fmt.Scanf("%d\n", &opt); opt {
		case 1:
			telaInformacaoUsuario()
			utils.Pause()
		case 2:
			telaInformacaoEventos()
			utils.Pause()
		}
	}
}

func telaInformacaoUsuario() {
	utils.ClearScreen()
	fmt.Println("Nao implementado ainda")
}

func telaInformacaoEventos() {
	utils.ClearScreen()
	fmt.Println("Nao implementado ainda")
}
