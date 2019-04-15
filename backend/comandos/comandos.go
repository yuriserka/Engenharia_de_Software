package comandos

import (
	ctrlservicos "github.com/Engenharia_de_Software/backend/controladoras_servicos"
	"github.com/Engenharia_de_Software/entidades"
)

func CadastrarNovoUsuario(usuario entidades.Usuario) {
	ctrlservicos.CadastrarUsr(usuario)
}
