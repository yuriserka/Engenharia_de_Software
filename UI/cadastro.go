package ui

import (
	"fmt"

	"github.com/yuriserka/Engenharia_de_Software/api/controladoras"
	"github.com/yuriserka/Engenharia_de_Software/api/utils"
)

// Cadastrar é responsável por pegar os dados do usuário e inserir
// no banco de dados
func cadastrarUsuario() {
	utils.ClearScreen()
	fmt.Println("\tCadastre-se")

	cpf, senha := utils.GetUserData()

	if ok := controladoras.CadastrarNovoUsuario(cpf, senha); ok {
		fmt.Println("Cadastrado com sucesso!")
	}
	utils.Pause()
}
