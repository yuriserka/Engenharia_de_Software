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
	repositorios.SetCartao(cpf, numCartao, codCartao, valCartao)
}

// MostrarCartoes mostra todos os cartoes que o usuario tem cadastrado
func MostrarCartoes(cpf string) {
	cartoes := repositorios.GetCartoesUsuario(cpf)
	for _, cartao := range cartoes {
		fmt.Printf("Numero: %s\nCódigo %s\nValidade: %s\n", cartao.Numero, cartao.Codigo, cartao.Validade)
		fmt.Println(strings.Repeat("-", 10))
	}
}

// MudarSenha atualiza a senha do usuario
func MudarSenha(cpf, senha string) {
	repositorios.UpdateUsuario(cpf, senha)
}
