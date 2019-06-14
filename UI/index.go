package ui

import (
	"fmt"

	"github.com/yuriserka/Engenharia_de_Software/api/controladoras"

	"github.com/yuriserka/Engenharia_de_Software/api/utils"
)

// Init apresentará as opções de Login e Cadastro. É onde o programa começa.
func Init() {
	var opt int
	for opt != 3 {
		utils.ClearScreen()
		fmt.Print("\t\tSistema de Vendas de Ingressos\n", "Escolha uma das opções abaixo:\n")
		fmt.Print("[1] login\n", "[2] cadastrar-se\n", "[3] sair\n", "\topção: ")

		switch fmt.Scanf("%d\n", &opt); opt {
		case 1:
			if cpf, logado := entrar(); logado {
				controleLogado(cpf)
			}
		case 2:
			cadastrar()
			utils.Pause()
		}
	}
}

// controleLogado apresentará as opções de gestão de Usuário e de Eventos. Só é acessível após o usuário ser
// autenticado.
func controleLogado(cpf string) {
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

// entrar solicita os dados do usuário e verifica se o mesmo está no bando de dados. Se estiver então é retornado
// o cpf dele e True, caso contrário retorna uma string vazia e False.
func entrar() (string, bool) {
	utils.ClearScreen()
	fmt.Println("\tEntrar")

	cpf, senha := utils.GetUserData()

	if logado := controladoras.Autenticar(cpf, string(senha)); !logado {
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
