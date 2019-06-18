package servicos

import (
		"testing"
)

func TestCadastrarNovoUsuario(t *testing.T){
	cpf := 12345678900
	senha := "1234"
	got := CadastrarNovoUsuario(cpf, senha)
	if got ≠ true {
		t.Errorf("CadastrarNovoUsuario(%d, %s) = %t; want true", cpf, senha, got)
	}

}
