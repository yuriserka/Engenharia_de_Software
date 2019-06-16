package ui

import (
	"fmt"

	"github.com/yuriserka/Engenharia_de_Software/api/controladoras"
	"github.com/yuriserka/Engenharia_de_Software/api/utils"
)

// entrar solicita os dados do usuário e verifica se o mesmo está no bando de dados. Se estiver então é retornado
// o cpf dele e True, caso contrário retorna uma string vazia e False.
func entrar() (string, bool) {
	utils.ClearScreen()
	fmt.Println("\tEntrar")

	cpf, senha := utils.GetUserData()

	if erro := controladoras.Autenticar(cpf, string(senha)); erro != nil {
		fmt.Println(erro.Error())
		utils.Pause()
		return "", false
	}

	return cpf, true
}
