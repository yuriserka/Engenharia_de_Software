package telas

import (
	"fmt"

	"github.com/Engenharia_de_Software/utils"
)

// TelaPrincipal é onde o main loop da aplicação ocorre
func TelaPrincipal() {
	var opt int
	for opt != 3 {
		utils.ClearScreen()
		fmt.Print("\t\tSistema de Vendas de Ingressos\n", "Escolha uma das opções abaixo:\n")
		fmt.Print("[1] login\n", "[2] cadastrar-se\n", "[3] sair\n", "\topção: ")
		switch fmt.Scanf("%d\n", &opt); {
		case opt == 1:
			telaLogin()
			utils.Pause()
		case opt == 2:
			telaCadastroUsuario()
			utils.Pause()
		}
	}
}
