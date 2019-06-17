package controladoras

import (
		"testing"
)

func TestRecuperarUsuario(t *testing.T){
	cpf := "01234567890"
	got := RecuperarUsuario(cpf)

	// Testa se o CPF retornado tem onze dígitos
	if len(cpf) != 11 {
		t.Errorf("[1] RecuperarUsuario(%s) = %v; want 11 char string", cpf, got)
	}
}

func TestCadastrarCartaoCredito(t *testing.T){
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

func TestRecuperarCartoesDeCredito(t *testing.T){
	cpf := "01234567890"
	got := RecuperarCartoesDeCredito(cpf)
	if len(got) == 0 {
		t.Errorf("[4] RecuperarCartoesDeCredito(%s) = %v; want map with a key-value pair", cpf, got)
	}
}
