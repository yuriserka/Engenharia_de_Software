package comandos

import (
	ctrlservicos "github.com/Engenharia_de_Software/backend/controladoras_servicos"
	"github.com/Engenharia_de_Software/entidades"
	"github.com/Engenharia_de_Software/utils"
)

// CadastrarNovoUsuario criptografa a senha e redireciona o fluxo de execução para a controladora de serviços.
func CadastrarNovoUsuario(usuario entidades.Usuario) {
	senhaCriptografada := utils.CriptografaSenha(usuario.Senha)
	ctrlservicos.CadastrarUsr(usuario.Cpf, senhaCriptografada)
}
