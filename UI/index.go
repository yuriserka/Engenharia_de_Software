package ui

import (
	"fmt"

	"github.com/yuriserka/Engenharia_de_Software/api/controladoras"

	"github.com/yuriserka/Engenharia_de_Software/api/utils"
)

// Init apresentará as opções de Login e Cadastro. É onde o programa começa.
func Init() {
	const (
		login       = 1
		cadastrarse = 2
		sair        = 3
	)
	menu := map[int]string{
		login:       "Fazer Login",
		cadastrarse: "Cadastrar-se",
		sair:        "Sair",
	}
	var opt int
	sortedIndexes := utils.OrdenaMap(menu)
	for opt != sair {
		utils.ClearScreen()
		fmt.Print("\tBem vindo Ao sistema de venda de ingressos.\n", "Escolha uma das opções abaixo:\n")

		for _, i := range sortedIndexes {
			fmt.Printf("[%d] %s\n", i, menu[i])
		}

		fmt.Print("\tOpcao: ")
		switch fmt.Scanf("%d\n", &opt); opt {
		case login:
			if cpf, logado := entrar(); logado {
				controleLogado(cpf)
			}
		case cadastrarse:
			cadastrar()
			utils.Pause()
		}
	}
}

// controleLogado apresentará as opções de gestão de Usuário e de Eventos. Só é acessível após o usuário ser
// autenticado.
func controleLogado(cpf string) {
	const (
		infoUsuario = 1
		infoEventos = 2
		voltar      = 3
	)
	menu := map[int]string{
		infoUsuario: "Informacoes do Usuario",
		infoEventos: "Visualizar Eventos",
		voltar:      "Voltar",
	}
	var opt int
	sortedIndexes := utils.OrdenaMap(menu)
	for opt != voltar {
		utils.ClearScreen()
		fmt.Print("\tBem vindo\n", "Escolha uma das opções abaixo:\n")

		for _, i := range sortedIndexes {
			fmt.Printf("[%d] %s\n", i, menu[i])
		}

		fmt.Print("\tOpcao: ")
		switch fmt.Scanf("%d\n", &opt); opt {
		case infoUsuario:
			gestaoUsuario(cpf)
		case infoEventos:
			gestaoEventos(cpf)
		}
	}
}

// entrar solicita os dados do usuário e verifica se o mesmo está no bando de dados. Se estiver então é retornado
// o cpf dele e True, caso contrário retorna uma string vazia e False.
func entrar() (string, bool) {
	utils.ClearScreen()
	fmt.Println("\tEntrar")

	cpf, senha := utils.GetUserData()

	if autenticado := controladoras.Autenticar(cpf, string(senha)); !autenticado {
		utils.Pause()
		return "", false
	}

	return cpf, true
}

// Cadastrar é responsável por pegar os dados do usuário e inserir
// no banco de dados
func cadastrar() {
	utils.ClearScreen()
	fmt.Println("\tCadastre-se")

	cpf, senha := utils.GetUserData()
	senhaCriptografada := utils.CriptografaSenha(senha)

	if ok := controladoras.CadastrarNovoUsuario(cpf, senhaCriptografada); ok {
		fmt.Println("Cadastrado com sucesso!")
	}
}

func gestaoUsuario(cpf string) {
	const (
		editarSenha       = 1
		cadastrarCartao   = 2
		visualizarCartoes = 3
		excluirConta      = 4
		voltar            = 5
	)
	menu := map[int]string{
		editarSenha:       "Editar Senha",
		cadastrarCartao:   "Cadastrar Cartao",
		visualizarCartoes: "Visualizar Cartoes",
		excluirConta:      "Excluir Conta",
		voltar:            "Voltar",
	}
	var opt int
	sortedIndexes := utils.OrdenaMap(menu)
	for opt != voltar {
		utils.ClearScreen()
		fmt.Println("\tPerfil")
		controladoras.MostrarUsuario(cpf)

		for _, i := range sortedIndexes {
			fmt.Printf("[%d] %s\n", i, menu[i])
		}

		fmt.Println("\tOpcao: ")
		switch fmt.Scanf("%d\n", &opt); opt {
		case editarSenha:
			utils.NaoImplementado("Editar Senha")
		case cadastrarCartao:
			cadastrarCartoes(cpf)
		case visualizarCartoes:
			verCartoes(cpf)
		case excluirConta:
			utils.NaoImplementado("Excluir Conta")
		}
	}
}

func verCartoes(cpf string) {
	utils.ClearScreen()
	fmt.Println("\tVisualizando Cartões de Crédito")
	controladoras.MostrarCartoes(cpf)
	utils.Pause()
}

func cadastrarCartoes(cpf string) {
	utils.ClearScreen()
	fmt.Println("\tCadastro de Cartão de Crédito")
	var (
		num, cod, val string
	)
	fmt.Printf("numero do cartão: ")
	fmt.Scanf("%s\n", &num)
	fmt.Printf("coddigo do cartão: ")
	fmt.Scanf("%s\n", &cod)
	fmt.Printf("validade do cartão: ")
	fmt.Scanf("%s\n", &val)

	controladoras.CadastrarCartaoCredito(cpf, num, cod, val)
	utils.Pause()
}

func gestaoEventos(cpf string) {
	utils.ClearScreen()
	fmt.Println("\tGestão dos Eventos")
}
