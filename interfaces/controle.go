package interfaces

// InterfaceServicoControle é a interface responsável pela integração das interfaces do sistema.
type InterfaceServicoControle interface {
	Build()
}

// InterfaceApresentacaoControle é a interface responsável pela tela inicial e após ser efetuado login.
// Assim como possui instâncias das interfaces de apresentação do sistema.
type InterfaceApresentacaoControle interface {
	Init()
	controleLogado(cpf string)
	SetControladoraApresentacaoLogin(InterfaceApresentacaoLogin)
	SetControladoraApresentacaoCadastro(InterfaceApresentacaoCadastro)
	SetControladoraApresentacaoUsuario(InterfaceApresentacaoUsuario)
	SetControladoraApresentacaoEventos(InterfaceApresentacaoEventos)
}
