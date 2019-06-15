package controladoras

import (
	"fmt"
	"strings"

	"github.com/yuriserka/Engenharia_de_Software/api/repositorios"
)

// CadastrarNovoEvento cadastra novos eventos e os associa ao usuario criador
func CadastrarNovoEvento(cpf, codigo, nome, cidade, estado, tipo, classificacao string) {
	repositorios.SetEvento(codigo, nome, cidade, estado, tipo, classificacao)
	repositorios.SetEventoUsuario(cpf, codigo)
}

// MostrarTodosEventos Mostra informação sobre todos os eventos cadastrados
func MostrarTodosEventos() {
	eventos := repositorios.GetTodosEventos()
	for _, evento := range eventos {
		fmt.Println(strings.Repeat("-", 10))
		fmt.Printf("Codigo: %s\nNome: %s\nLocal: %s\nClassificação: %s\nTipo: %s\n", evento.Codigo,
			evento.Nome, evento.Cidade+"-"+evento.Estado, evento.Classificacao, evento.Tipo)
	}
	fmt.Println()
}
