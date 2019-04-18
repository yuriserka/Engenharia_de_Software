package telas

import (
	"fmt"

	"github.com/Engenharia_de_Software/backend/comandos"
	ctrlServicos "github.com/Engenharia_de_Software/backend/controladoras_servicos"
	"github.com/Engenharia_de_Software/utils"
)

func telaLogin() {
	utils.ClearScreen()
	fmt.Println("\tEntrar")

	loginInterface := comandos.InterfaceApresentacaoLogin{}
	servico := ctrlServicos.InterfaceServicoLogin{}

	loginInterface.Executar(&servico)
}
