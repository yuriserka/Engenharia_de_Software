package repositorios

import (
		"testing"
		"github.com/yuriserka/Engenharia_de_Software/api/controladoras"
		"github.com/yuriserka/Engenharia_de_Software/api/entidades"
)

func TestSetApresentacaoEvento(t *testing.T){
	cpf := "01234567890"
	codigo := "TestSetApresentacaoEvento"
	nome := "evento"
	cidade := "cidade"
	estado := "estado"
	tipo := "tipo"
	classificacao := "classificacao"
	controladoras.CadastrarNovoEvento(cpf, codigo, nome, cidade, estado, tipo, classificacao)

	// Testa set de nova apresentaçao, deve retornar nil no sucesso
	got := SetApresentacaoEvento(codigo, &entidades.Apresentacao{codigo, "sala1", 10, "hoje", "agora", 10.43})
	if got != nil {
		t.Errorf("SetApresentacaoEvento(%s, &entidades.Apresentacao{%s, \"sala1\", 10, \"hoje\", \"agora\", 10.43}) = %v; want nil", codigo, codigo, got)
	}

	// Testa set de apresentaçao já setada, deve retornar erro na falha
	got = SetApresentacaoEvento(codigo, &entidades.Apresentacao{codigo, "sala1", 10, "hoje", "agora", 10.43})
	if got == nil {
		t.Errorf("SetApresentacaoEvento(%s, &entidades.Apresentacao{%s, \"sala1\", 10, \"hoje\", \"agora\", 10.43}) = %v; want error", codigo, codigo, got)
	}

	// TODO: Fazer um teste para testar dez apresentações num evento
}

func TestGetApresentacoes(t *testing.T){
	codigo := "TestSetApresentacaoEvento"

	_,got := GetApresentacoes(codigo)
	if got != nil {
		t.Errorf("GetApresentacoes(%s) = %v; want (_, nil)", codigo, got)
	}

	codigo = "TestGetApresentacoes"
	_,got = GetApresentacoes(codigo)
	if got == nil {
		t.Errorf("GetApresentacoes(%s) = %v; want (nil, error)", codigo, got)
	}
}







