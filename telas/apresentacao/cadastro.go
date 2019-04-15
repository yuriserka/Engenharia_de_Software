package telas

import (
	"fmt"

	telasServico "github.com/Engenharia_de_Software/telas/servico"
	"github.com/Engenharia_de_Software/utils"
)

func telaCadastroUsuario() {
	utils.ClearScreen()
	fmt.Println("\tCadastro de Usuarios")

	telasServico.CadastrarUsuario()
}
