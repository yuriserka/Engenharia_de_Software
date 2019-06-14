package utils

import (
	"fmt"

	"github.com/howeyc/gopass"
)

// GetUserData recebe o input do usuario tanto no login quanto no cadastro
func GetUserData() (cpf string, senha []byte) {
	fmt.Printf("Cpf: ")
	fmt.Scanf("%s\n", &cpf)

	fmt.Printf("Senha: ")
	senha, err := gopass.GetPasswdMasked()

	if err != nil {
		panic(err.Error())
	}

	return
}
