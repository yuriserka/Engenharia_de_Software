package ui

import (
	"fmt"

	"github.com/yuriserka/Engenharia_de_Software/api/controladoras"
	"github.com/yuriserka/Engenharia_de_Software/api/utils"
)

func gestaoUsuario(cpf string) {
	const (
		keditarSenha       = iota + 1
		kcadastrarCartao   = iota + 1
		kvisualizarCartoes = iota + 1
		kvisualizarEventos = iota + 1
		kexcluirConta      = iota + 1
		kvoltar            = iota + 1
	)
	menu := map[int]string{
		keditarSenha:       "Editar Senha",
		kcadastrarCartao:   "Cadastrar Cartao",
		kvisualizarCartoes: "Visualizar Cartoes",
		kvisualizarEventos: "Visualizar Eventos Criados",
		kexcluirConta:      "Excluir Conta",
		kvoltar:            "Voltar",
	}
	var opt int
	sortedIndexes := utils.OrdenaMap(menu)
	for opt != kvoltar {
		utils.ClearScreen()
		fmt.Println("\tPerfil")
		controladoras.MostrarUsuario(cpf)

		for _, i := range sortedIndexes {
			fmt.Printf("[%d] %s\n", i, menu[i])
		}

		fmt.Print("\tOpcao: ")
		switch fmt.Scanf("%d\n", &opt); opt {
		case keditarSenha:
			editarSenha(cpf)
		case kcadastrarCartao:
			cadastrarCartoes(cpf)
		case kvisualizarCartoes:
			visualizarCartoes(cpf)
		case kvisualizarEventos:
			visualizarEventos(cpf)
		case kexcluirConta:
			utils.NaoImplementado("Excluir Conta")
		}
	}
}

func editarSenha(cpf string) {
	utils.ClearScreen()
	fmt.Println("\tEdite sua Senha")
	var novaSenha string
	fmt.Printf("Digite a nova senha: ")
	fmt.Scanf("%s\n", &novaSenha)
	controladoras.MudarSenha(cpf, novaSenha)
	utils.Pause()
}

func visualizarCartoes(cpf string) {
	utils.ClearScreen()
	fmt.Printf("\tVisualizando Cartões de Crédito\n\n")
	controladoras.MostrarCartoes(cpf)
	utils.Pause()
}

func cadastrarCartoes(cpf string) {
	utils.ClearScreen()
	fmt.Println("\tCadastro de Cartão de Crédito")
	var (
		num, cod, val string
	)
	fmt.Printf("numero: ")
	fmt.Scanf("%s\n", &num)
	fmt.Printf("coddigo: ")
	fmt.Scanf("%s\n", &cod)
	fmt.Printf("validade: ")
	fmt.Scanf("%s\n", &val)

	controladoras.CadastrarCartaoCredito(cpf, num, cod, val)
	utils.Pause()
}

func visualizarEventos(cpf string) {
	utils.ClearScreen()
	fmt.Printf("\tVisualizando Eventos Criados\n\n")
	controladoras.MostrarEventosUsuario(cpf)
	utils.Pause()
}
