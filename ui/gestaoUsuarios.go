package ui

import (
	"fmt"
	"strings"

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
		fmt.Println(controladoras.RecuperarCpfFormatado(cpf))

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

	fmt.Println("Sucesso ao mudar senha")
	utils.Pause()
}

func visualizarCartoes(cpf string) {
	utils.ClearScreen()
	fmt.Printf("\tVisualizando Cartões de Crédito\n\n")

	ccs, e1 := controladoras.RecuperarCartoesDeCredito(cpf)
	if e1 != nil {
		fmt.Println(e1.Error())
		utils.Pause()
		return
	}
	for _, cartao := range ccs {
		fmt.Println(strings.Repeat("-", 10))
		fmt.Print(cartao.ToString())
	}
	fmt.Println()

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

	if erro := controladoras.CadastrarCartaoCredito(cpf, num, cod, val); erro != nil {
		fmt.Println(erro.Error())
	} else {
		fmt.Println("Cartão cadastrado com sucesso")
	}

	utils.Pause()
}

func visualizarEventos(cpf string) {
	utils.ClearScreen()
	fmt.Printf("\tVisualizando Eventos Criados\n\n")

	eventos, e1 := controladoras.RecuperarEventosUsuario(cpf)
	if e1 != nil {
		fmt.Println(e1.Error())
		utils.Pause()
		return
	}
	for _, e := range eventos {
		fmt.Println(strings.Repeat("-", 10))
		fmt.Printf("Codigo: %s\nNome: %s\nLocal: %s\nClassificação: %s\nTipo: %s\n", e.Codigo,
			e.Nome, e.Cidade+"-"+e.Estado, e.Classificacao, e.Tipo)
	}
	fmt.Println()

	utils.Pause()
}
