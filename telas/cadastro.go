package telas

import (
	"fmt"

	"github.com/Engenharia_de_Software/backend/comandos"

	ctrlServicos "github.com/Engenharia_de_Software/backend/controladoras_servicos"

	"github.com/Engenharia_de_Software/utils"
)

func telaCadastroUsuario() {
	utils.ClearScreen()
	fmt.Println("\tCadastre-se")

	registerInterface := comandos.InterfaceApresentacaoCadastro{}
	servico := ctrlServicos.InterfaceServicoCadastro{}

	registerInterface.Executar(&servico)
}
