package controladoras

import (
	"strings"

	"github.com/yuriserka/Engenharia_de_Software/api/entidades"

	"github.com/yuriserka/Engenharia_de_Software/api/repositorios"
)

// RecuperarUsuario mostra o CPF do usuario de forma bonita,
// se não tiver ponto ele imprime do mesmo jeito, mas errado
func RecuperarUsuario(cpf string) string {
	u := repositorios.GetUsuario(cpf)
	splitAfterN := func(n int) (substr []string) {
		a := []rune(u.Cpf)
		var res string
		for i, r := range a {
			res = res + string(r)
			if i == len(cpf)-1 {
				substr = append(substr, res)
				res = ""
			} else if i > 0 && (i+1)%n == 0 {
				substr = append(substr, res)
				res = ""
			}
		}
		return
	}
	cpfComPonto := strings.Join(splitAfterN(3), ".")
	idx := strings.LastIndex(cpfComPonto, ".")
	if idx == -1 {
		return cpfComPonto
	}
	cpfCorreto := cpfComPonto[:idx] + strings.Replace(cpfComPonto[idx:], ".", "-", 1)
	return cpfCorreto
}

// CadastrarCartaoCredito Adiciona um cartao para o usuario
func CadastrarCartaoCredito(cpf, numCartao, codCartao, valCartao string) error {
	return repositorios.SetCartao(cpf, numCartao, codCartao, valCartao)
}

// RecuperarCartoesDeCredito mostra todos os cartoes que o usuario tem cadastrado
func RecuperarCartoesDeCredito(cpf string) map[string]*entidades.CartaoDeCredito {
	return repositorios.GetCartoesUsuario(cpf)
}

// MudarSenha atualiza a senha do usuario
func MudarSenha(cpf, senha string) {
	repositorios.UpdateUsuario(cpf, senha)
}

// RecuperarEventosUsuario imprime os eventos que o usuario criou
func RecuperarEventosUsuario(cpf string) []*entidades.Evento {
	codEventos := repositorios.GetEventosUsuario(cpf)
	eventos := make([]*entidades.Evento, 0, len(codEventos))
	for _, cod := range codEventos {
		eventos = append(eventos, repositorios.GetEvento(cod))
	}
	return eventos
}
