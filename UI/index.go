package ui

import (
	"fmt"

	"github.com/yuriserka/Engenharia_de_Software/api/controladoras"

	"github.com/yuriserka/Engenharia_de_Software/api/utils"
)

// Init apresentará as opções de Login e Cadastro. É onde o programa começa.
func Init() {
	const (
		klogin     = 1
		kcadastrar = 2
		ksair      = 3
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
			cadastrar()
			utils.Pause()
		}
	}
}

// controleLogado apresentará as opções de gestão de Usuário e de Eventos. Só é acessível após o usuário ser
// autenticado.
func controleLogado(cpf string) {
	const (
		kinfoUsuario = 1
		kinfoEventos = 2
		kvoltar      = 3
	)
	menu := map[int]string{
		kinfoUsuario: "Informacoes do Usuario",
		kinfoEventos: "Visualizar Eventos",
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

	if ok := controladoras.CadastrarNovoUsuario(cpf, senha); ok {
		fmt.Println("Cadastrado com sucesso!")
	}
}

func gestaoUsuario(cpf string) {
	const (
		keditarSenha       = 1
		kcadastrarCartao   = 2
		kvisualizarCartoes = 3
		kexcluirConta      = 4
		kvoltar            = 5
	)
	menu := map[int]string{
		keditarSenha:       "Editar Senha",
		kcadastrarCartao:   "Cadastrar Cartao",
		kvisualizarCartoes: "Visualizar Cartoes",
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
	utils.NaoImplementado("Gestao de Eventos")
}
