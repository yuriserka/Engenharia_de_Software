package ctrlservicos

import (
	"fmt"
	"time"

	"github.com/Engenharia_de_Software/entidades"
)

// CadastrarUsr registra o novo usuario no banco de dados.
func CadastrarUsr(usr entidades.Usuario) {
	// conectar no banco de dados e blabla
	fmt.Println("inserindo no banco de dados")
	time.Sleep(2)
	fmt.Println("Tudo certo, usuario cadastrado")
}
