package interfaces

// InterfaceServicoLogin é a interface responsável pelo login dos usuários do sistema.
type InterfaceServicoLogin interface {
	Autenticar(cpf, senha string) bool
}

// InterfaceApresentacaoLogin é a interface responsável pela solicitação dos dados de um usuário para fazer login no sistema.
type InterfaceApresentacaoLogin interface {
	Autenticar() (string, bool)
	SetControladoraServicoLogin(InterfaceServicoLogin)
}
