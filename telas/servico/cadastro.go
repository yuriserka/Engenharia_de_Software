package telas

import (
	"fmt"

	"github.com/Engenharia_de_Software/backend/comandos"
	"github.com/Engenharia_de_Software/entidades"

	"github.com/howeyc/gopass"
)

func CadastrarUsuario() {
	var cpf string
	fmt.Printf("Cpf: ")
	fmt.Scanf("%s\n", &cpf)
	fmt.Printf("Senha: ")
	senha, err := gopass.GetPasswdMasked()
	if err != nil {
		panic(err.Error())
	}

	comandos.CadastrarNovoUsuario(entidades.Usuario{
		Cpf:   cpf,
		Senha: string(senha),
	})
}
