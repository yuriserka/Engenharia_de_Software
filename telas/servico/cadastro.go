package telas

import (
	"fmt"

	comandos "github.com/Engenharia_de_Software/backend/comandos"
	"github.com/Engenharia_de_Software/entidades"
)

func CadastrarUsuario() {
	usr := entidades.Usuario{}

	fmt.Println("Digite seu CPF")
	fmt.Scanln(&usr.Cpf)
	fmt.Println("Digite sua senha")
	fmt.Scanln(&usr.Senha)

	comandos.CadastrarNovoUsuario(usr)
}
