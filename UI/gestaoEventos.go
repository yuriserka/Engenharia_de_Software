package ui

import (
	"bufio"
	"fmt"
	"os"

	"github.com/yuriserka/Engenharia_de_Software/api/controladoras"

	"github.com/yuriserka/Engenharia_de_Software/api/utils"
)

func gestaoEventos(cpf string) {
	const (
		kcadastrar  = iota + 1
		kvisualizar = iota + 1
		kvoltar     = iota + 1
	)
	menu := map[int]string{
		kcadastrar:  "Cadastrar Evento",
		kvisualizar: "Visualizar Todos os Eventos",
		kvoltar:     "Voltar",
	}
	var opt int
	sortedIndexes := utils.OrdenaMap(menu)
	for opt != kvoltar {
		utils.ClearScreen()
		fmt.Println("\tGestão de Eventos")

		for _, i := range sortedIndexes {
			fmt.Printf("[%d] %s\n", i, menu[i])
		}

		fmt.Print("\tOpcao: ")
		switch fmt.Scanf("%d\n", &opt); opt {
		case kcadastrar:
			cadastrarEvento(cpf)
		case kvisualizar:
			visualizarTodosEventos()
		}
	}
}

func cadastrarEvento(cpf string) {
	utils.ClearScreen()
	fmt.Println("Cadastro de Evento")

	reader := bufio.NewReader(os.Stdin)

	removerNovaLinhaInput := func(s string, err error) string {
		s = s[:len(s)-2]
		return s
	}

	fmt.Printf("Codigo: ")
	cod := removerNovaLinhaInput(reader.ReadString('\n'))
	fmt.Printf("Nome: ")
	nome := removerNovaLinhaInput(reader.ReadString('\n'))
	fmt.Printf("Estado: ")
	estado := removerNovaLinhaInput(reader.ReadString('\n'))
	fmt.Printf("Cidade: ")
	cidade := removerNovaLinhaInput(reader.ReadString('\n'))
	fmt.Printf("Tipo: ")
	tipo := removerNovaLinhaInput(reader.ReadString('\n'))
	fmt.Printf("Classificação: ")
	classificacao := removerNovaLinhaInput(reader.ReadString('\n'))

	controladoras.CadastrarNovoEvento(cpf, cod, nome, cidade, estado, tipo, classificacao)
	utils.Pause()
}

func visualizarTodosEventos() {
	utils.ClearScreen()
	fmt.Printf("\tVisualizando Todos so Eventos\n\n")
	controladoras.MostrarTodosEventos()
	utils.Pause()
}
