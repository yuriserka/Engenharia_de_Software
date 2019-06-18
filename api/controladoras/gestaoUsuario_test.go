package controladoras

import (
	"testing"
)

func TestRecuperarCPFCompletoFormatado(t *testing.T) {
	cpf := "01234567890"
	got := RecuperarCpfFormatado(cpf)

	// Testa se o CPF retornado tem 14 dígitos
	if len(cpf) != 14 {
		t.Errorf("[1] RecuperarCpfFormatado(%s) = %v; want 14 char string", cpf, got)
	}
}

func TestRecuperarUsuario(t *testing.T) {

}

func TestCadastrarCartaoCredito(t *testing.T) {
	cpf := "01234567890"
	numCartao := "1"
	codCartao := "2"
	valCartao := "01/21"

	// Testa cadastro de novo cartão, deve retornar nil no sucesso
	got := CadastrarCartaoCredito(cpf, numCartao, codCartao, valCartao)
	if got != nil {
		t.Errorf("[2] CadastrarCartaoCredito(%s, %s, %s, %s) = %v; want nil", cpf, numCartao, codCartao, valCartao, got)
	}

	// Testa cadastro de usuário já cadastrado, deve retornar erro na falha
	got = CadastrarCartaoCredito(cpf, numCartao, codCartao, valCartao)
	if got == nil {
		t.Errorf("[3] CadastrarCartaoCredito(%s, %s, %s, %s) = %v; want error", cpf, numCartao, codCartao, valCartao, got)
	}
}

func TestRecuperarCartoesDeCredito(t *testing.T) {
	cpf := "01234567890"
	_, got := RecuperarCartoesDeCredito(cpf)
	if got == nil {
		t.Errorf("[4] RecuperarCartoesDeCredito(%s) = %v; want map with a key-value pair", cpf, got)
	}
}
