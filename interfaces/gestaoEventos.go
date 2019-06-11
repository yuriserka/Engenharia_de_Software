package interfaces

// InterfaceServicoEventos
type InterfaceServicoEventos interface {
	// tudo que será necessario na parte de eventos
	CriarEvento()
}

// InterfaceApresentacaoEventos
type InterfaceApresentacaoEventos interface {
	Executar(cpf string)
	SetControladoraServicoEventos(InterfaceServicoEventos)
}
