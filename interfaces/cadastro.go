package interfaces

// InterfaceServicoCadastro é a interface responsável pelo cadastro de um novo usuário ao sistema.
type InterfaceServicoCadastro interface {
	CadastrarNovoUsuario(cpf, senha string) bool
}

// InterfaceApresentacaoCadastro é a interface responsável pela solicitação dos dados do novo usuário do sistema.
type InterfaceApresentacaoCadastro interface {
	Cadastrar()
	SetControladoraServicoCadastro(InterfaceServicoCadastro)
}
