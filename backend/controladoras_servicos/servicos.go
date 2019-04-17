package ctrlservicos

import (
	"fmt"
	"time"
)

// CadastrarUsr registra um usuário no banco de dados.
func CadastrarUsr(cpf, senha string) {
	// conectar no banco de dados e blabla
	fmt.Printf("inserindo usuario {Cpf: %s, Senha: %s} no banco de dados\n", cpf, senha)
	time.Sleep(2 * time.Second)
	fmt.Println("Tudo certo, usuario cadastrado")
}
