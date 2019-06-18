package controladoras

import (
	"testing"
)

func TestCadastrarNovoEvento(t *testing.T) {
	cpf := "01234567890"
	codigo := "1"
	nome := "Eventao"
	cidade := "Brasília"
	estado := "DF"
	tipo := "Legal"
	classificacao := "12+"

	// Testa cadastro de novo evento, deve retornar nil no sucesso
	got := CadastrarNovoEvento(cpf, codigo, nome, cidade, estado, tipo, classificacao)
	if got != nil {
		t.Errorf("[1] CadastrarNovoEvento(%s, %v, %s, %s, %s, %s, %s) = %v; want nil", cpf, codigo, nome, cidade, estado, tipo, classificacao, got)
	}

	// Testa cadastro de evento já cadastrado, deve retornar erro na falha
	got = CadastrarNovoEvento(cpf, codigo, nome, cidade, estado, tipo, classificacao)
	if got == nil {
		t.Errorf("[2] CadastrarNovoEvento(%s, %v, %s, %s, %s, %s, %s) = %v; want error", cpf, codigo, nome, cidade, estado, tipo, classificacao, got)
	}

	// Testa cadasto de evento com o mesmo código, deve retornar erro na falha
	nome1 := "Evento2"
	got = CadastrarNovoEvento(cpf, codigo, nome1, cidade, estado, tipo, classificacao)
	if got == nil {
		t.Errorf("[3] CadastrarNovoEvento(%s, %v, %s, %s, %s, %s, %s) = %v; want error", cpf, codigo, nome1, cidade, estado, tipo, classificacao, got)
	}
}
