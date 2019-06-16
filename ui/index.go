package ui

import (
	"fmt"

	"github.com/yuriserka/Engenharia_de_Software/api/utils"
)

// Init apresentará as opções de Login e Cadastro. É onde o programa começa.
func Init() {
	const (
		klogin     = iota + 1
		kcadastrar = iota + 1
		ksair      = iota + 1
	)
	menu := map[int]string{
		klogin:     "Fazer Login",
		kcadastrar: "Cadastrar-se",
		ksair:      "Sair",
	}
	var opt int
	sortedIndexes := utils.OrdenaMap(menu)
	for opt != ksair {
		utils.ClearScreen()
		fmt.Print("\tBem vindo Ao sistema de venda de ingressos.\n", "Escolha uma das opções abaixo:\n")

		for _, i := range sortedIndexes {
			fmt.Printf("[%d] %s\n", i, menu[i])
		}

		fmt.Print("\tOpcao: ")
		switch fmt.Scanf("%d\n", &opt); opt {
		case klogin:
			if cpf, logado := entrar(); logado {
				controleLogado(cpf)
			}
		case kcadastrar:
			cadastrarUsuario()
		}
	}
}

// controleLogado apresentará as opções de gestão de Usuário e de Eventos. Só é acessível após o usuário ser
// autenticado.
func controleLogado(cpf string) {
	const (
		kinfoUsuario = iota + 1
		kinfoEventos = iota + 1
		kvoltar      = iota + 1
	)
	menu := map[int]string{
		kinfoUsuario: "Gestão de Usuario",
		kinfoEventos: "Gestão de Eventos",
		kvoltar:      "Voltar",
	}
	var opt int
	sortedIndexes := utils.OrdenaMap(menu)
	for opt != kvoltar {
		utils.ClearScreen()
		fmt.Print("\tBem vindo\n", "Escolha uma das opções abaixo:\n")

		for _, i := range sortedIndexes {
			fmt.Printf("[%d] %s\n", i, menu[i])
		}

		fmt.Print("\tOpcao: ")
		switch fmt.Scanf("%d\n", &opt); opt {
		case kinfoUsuario:
			gestaoUsuario(cpf)
		case kinfoEventos:
			gestaoEventos(cpf)
		}
	}
}
