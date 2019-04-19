package interfaces

import (
	"github.com/Engenharia_de_Software/entidades"
)

// InterfaceServicoUsuario
type InterfaceServicoUsuario interface {
	Exibir(entidades.Usuario) (entidades.Usuario, bool)
	Editar(entidades.Usuario) (entidades.Usuario, bool)
	Excluir(cpf string) bool
}

// InterfaceApresentacaoUsuario
type InterfaceApresentacaoUsuario interface {
	Executar(cpf string) bool
	SetControladoraServicoUsuario(InterfaceServicoUsuario)
}
