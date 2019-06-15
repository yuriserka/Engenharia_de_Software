package controladoras

import (
	"fmt"
	"strings"

	"github.com/yuriserka/Engenharia_de_Software/api/repositorios"
)

// MostrarUsuario mostra o CPF do usuario de forma bonita,
// se não tiver ponto ele imprime do mesmo jeito, mas errado
func MostrarUsuario(cpf string) {
	u := repositorios.GetUsuario(cpf)
	splitAfterN := func() (substr []string) {
		a := []rune(u.Cpf)
		var res string
		for i, r := range a {
			res = res + string(r)
			if i == len(cpf)-1 {
				substr = append(substr, res)
				res = ""
			} else if i > 0 && (i+1)%3 == 0 {
				substr = append(substr, res)
				res = ""
			}
		}
		return
	}
	cpfComPonto := strings.Join(splitAfterN(), ".")
	idx := strings.LastIndex(cpfComPonto, ".")
	if idx == -1 {
		fmt.Println("CPF:", cpfComPonto)
		return
	}
	cpfCorreto := cpfComPonto[:idx] + strings.Replace(cpfComPonto[idx:], ".", "-", 1)
	fmt.Println("CPF:", cpfCorreto)
}

// CadastrarCartaoCredito Adiciona um cartao para o usuario
func CadastrarCartaoCredito(cpf, numCartao, codCartao, valCartao string) {
	erro := repositorios.SetCartao(cpf, numCartao, codCartao, valCartao)
	if erro != nil {
		fmt.Println(erro.Error())
	}
}

// MostrarCartoes mostra todos os cartoes que o usuario tem cadastrado
func MostrarCartoes(cpf string) {
	cartoes := repositorios.GetCartoesUsuario(cpf)
	for _, cartao := range cartoes {
		fmt.Println(strings.Repeat("-", 10))
		fmt.Printf("Numero: %s\nCódigo %s\nValidade: %s\n", cartao.Numero, cartao.Codigo, cartao.Validade)
	}
	fmt.Println()
}

// MudarSenha atualiza a senha do usuario
func MudarSenha(cpf, senha string) {
	repositorios.UpdateUsuario(cpf, senha)
}

// MostrarEventosUsuario imprime os eventos que o usuario criou
func MostrarEventosUsuario(cpf string) {
	eventos := repositorios.GetEventosUsuario(cpf)
	for _, codEvento := range eventos {
		fmt.Println(strings.Repeat("-", 10))
		evento := repositorios.GetEvento(codEvento)
		fmt.Printf("Codigo: %s\nNome: %s\nLocal: %s\nClassificação: %s\nTipo: %s\n", evento.Codigo,
			evento.Nome, evento.Cidade+"-"+evento.Estado, evento.Classificacao, evento.Tipo)
	}
	fmt.Println()
}
