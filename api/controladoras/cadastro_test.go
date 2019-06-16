package controladoras

import (
		"testing"
)

func TestCadastrarNovoUsuario(t *testing.T){
	cpf := "123456789"
	senha := []byte{1,2,3,4,5}

	// Testa cadastro de novo usuário, deve retornar nil no sucesso do cadastro
	got := CadastrarNovoUsuario(cpf, senha)
	if got != nil {
		t.Errorf("[1] CadastrarNovoUsuario(%s, %v) = %v; want nil", cpf, senha, got)
	}

	// Testa cadastro de usuário já cadastrado, deve retornar erro na falha do cadastro
	got = CadastrarNovoUsuario(cpf, senha)
	if got == nil {
		t.Errorf("[2] CadastrarNovoUsuario(%s, %v) = %v; want error", cpf, senha, got)
	}

	// Testa cadastro de outro usuário novo após já ter um cadastro, deve retornar nil no sucesso
	cpf = "987654321"
	got = CadastrarNovoUsuario(cpf, senha)
	if got != nil {
		t.Errorf("[3] CadastrarNovoUsuario(%s, %v) = %v; want error", cpf, senha, got)
	}

	// Testa cadastro de usuário com CPF vazio, deve retornar erro na falha
	cpf = ""
	got = CadastrarNovoUsuario(cpf, senha)
	if got == nil {
		t.Errorf("[4] CadastrarNovoUsuario(%s, %v) = %v; want error", cpf, senha, got)
	}
}
