package controladoras

import (
	"testing"
)

func TestAutenticar(t *testing.T) {
	cpf := "00000000000"
	senha := "senha"
	senhab := []byte{'s', 'e', 'n', 'h', 'a'}
	CadastrarNovoUsuario(cpf, senhab)

	// Testa autenticar um usuário cadastrado, retorna nil no sucesso
	got := Autenticar(cpf, senha)
	if got != nil {
		t.Errorf("[1] Autenticar(%s, %s) = %v; want nil", cpf, senha, got)
	}

	// Testa autenticar um usuário com senha errada, retorna erro na falha
	senha = "outrasenha"
	got = Autenticar(cpf, senha)
	if got == nil {
		t.Errorf("[2] Autenticar(%s, %s) = %v; want error (senha incorreta)", cpf, senha, got)
	}

	// Testa autenticar um usuário não cadastrado, retorna erro na falha
	cpf = "11111111111"
	got = Autenticar(cpf, senha)
	if got == nil {
		t.Errorf("[3] Autenticar(%s, %s) = %v; want error (usuario nao cadastrado ainda)", cpf, senha, got)
	}
}
