package ctrlservicos

import (
	"fmt"
	"time"

	"github.com/Engenharia_de_Software/entidades"
)

// CadastrarUsr registra o novo usuario no banco de dados.
func CadastrarUsr(usr entidades.Usuario) {
	// conectar no banco de dados e blabla
	fmt.Printf("inserindo usuario %v no banco de dados\n", usr)
	time.Sleep(2 * time.Second)
	fmt.Println("Tudo certo, usuario cadastrado")
}
