package servicos

import (
	"github.com/Engenharia_de_Software/telas"
)

// ControladoraServicoControle é responsável por setar as variáveis necessárias para o sistema funcionar.
//
// Implementa a InterfaceServicoControle.
type ControladoraServicoControle struct{}

// Build inicializa as interfaces do sistema.
func (ctrlSControle *ControladoraServicoControle) Build() {
	// sqlInit := new(backend.CriarTabelas)
	// sqlInit.Excutar()

	ctrlAL := &telas.ControladoraApresentacaoLogin{}
	ctrlAK := &telas.ControladoraApresentacaoCadastro{}
	// ctrlAU := new(telas.InterfaceApresentacaoUsuario)
	// ctrlAE := new(telas.InterfaceApresentacaoEventos)

	ctrlSL := &ControladoraServicoLogin{}
	ctrlSK := &ControladoraServicoCadastro{}

	// ctrlSU := new(InterfaceServicoUsuario)
	// ctrlSE := new(InterfaceServicoEventos)

	ctrlAL.SetControladoraServicoLogin(ctrlSL)
	ctrlAK.SetControladoraServicoCadastro(ctrlSK)
	// ctrlAU.SetControladoraServicoUsuario(ctrlSU)
	// ctrlAE.SetControladoraServicoEventos(ctrlSE)

	ctrlAC := &telas.ControladoraApresentacaoControle{}

	ctrlAC.SetControladoraApresentacaoLogin(ctrlAL)
	ctrlAC.SetControladoraApresentacaoCadastro(ctrlAK)
	// ctrlAC.SetControladoraApresentacaoUsuario(ctrlAU)
	// ctrlAC.SetControladoraApresentacaoEventos(ctrlAE)

	ctrlAC.Init()
}
