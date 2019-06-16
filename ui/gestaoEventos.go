package ui

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/yuriserka/Engenharia_de_Software/api/controladoras"

	"github.com/yuriserka/Engenharia_de_Software/api/utils"
)

func gestaoEventos(cpf string) {
	const (
		kcadastrarEvento         = iota + 1
		kcadastrarApresentacao   = iota + 1
		kvisualizarEventos       = iota + 1
		kvisualizarApresentacoes = iota + 1
		kvoltar                  = iota + 1
	)
	menu := map[int]string{
		kcadastrarEvento:         "Cadastrar Evento",
		kcadastrarApresentacao:   "Cadastrar Apresentação",
		kvisualizarEventos:       "Visualizar Todos os Eventos",
		kvisualizarApresentacoes: "Visualizar Apresentações de um Evento",
		kvoltar:                  "Voltar",
	}
	var opt int
	sortedIndexes := utils.OrdenaMap(menu)
	for opt != kvoltar {
		utils.ClearScreen()
		fmt.Printf("\tGestão de Eventos\n\n")

		for _, i := range sortedIndexes {
			fmt.Printf("[%d] %s\n", i, menu[i])
		}

		fmt.Print("\tOpcao: ")
		switch fmt.Scanf("%d\n", &opt); opt {
		case kcadastrarEvento:
			cadastrarEvento(cpf)
		case kvisualizarEventos:
			visualizarTodosEventos()
			utils.Pause()
		case kcadastrarApresentacao:
			cadastrarApresentacao()
		case kvisualizarApresentacoes:
			visualizarApresentacoes()
		}
	}
}

func cadastrarEvento(cpf string) {
	utils.ClearScreen()
	fmt.Printf("\tCadastro de Evento\n\n")

	v := struct {
		Cod, Nome, Estado, Cidade, Tipo, Classificacao string
	}{}

	utils.GetNStringInputs(&v, []string{
		"Código",
		"Nome",
		"Estado",
		"Cidade",
		"Tipo",
		"Classificação",
	}, 6)

	erro := controladoras.CadastrarNovoEvento(cpf, v.Cod, v.Nome, v.Cidade, v.Estado, v.Tipo, v.Classificacao)
	if erro != nil {
		fmt.Println(erro.Error())
	} else {
		fmt.Println("Evento criado com sucesso!")
	}
	utils.Pause()
}

func visualizarTodosEventos() {
	utils.ClearScreen()
	fmt.Printf("\tVisualizando Todos so Eventos\n\n")

	eventos := controladoras.RecuperarEventos()
	if len(eventos) == 0 {
		fmt.Println("Ainda não há eventos cadastrados")
		utils.Pause()
		return
	}
	for _, e := range eventos {
		fmt.Println(strings.Repeat("-", 10))
		fmt.Println(e.ToString())
	}
	fmt.Println()
}

func cadastrarApresentacao() {
	visualizarTodosEventos()

	fmt.Printf("\tEscolha um Evento pelo Código\n\n")
	var cod string
	fmt.Printf("codigo: ")
	fmt.Scanf("%s\n", &cod)

	e := controladoras.RecuperarEvento(cod)
	if e == nil {
		fmt.Println("Escolha uma opção válida")
		utils.Pause()
		return
	}

	utils.ClearScreen()

	v := struct {
		Codigo, Sala, Disponibilidade, Data, Horario, Preco string
	}{}

	utils.GetNStringInputs(&v, []string{
		"Código",
		"Sala",
		"Disponibilidade",
		"Data",
		"Horário",
		"Preço",
	}, 6)

	disp, _ := strconv.Atoi(v.Disponibilidade)
	pr, _ := strconv.ParseFloat(v.Preco, 64)
	erro := controladoras.CadastrarNovaApresentacao(v.Codigo, v.Sala, disp, v.Data, v.Horario, pr)
	if erro != nil {
		fmt.Println(erro.Error())
	} else {
		a := controladoras.RecuperarApresentacao(v.Codigo)
		if err := controladoras.LigarApresentacaoEvento(cod, a); err != nil {
			fmt.Println(err.Error())
		}
		fmt.Println("Apresentação cadastrada com sucesso")
	}

	utils.Pause()
}

func visualizarApresentacoes() {
	visualizarTodosEventos()

	fmt.Printf("\tEscolha um Evento pelo Código\n\n")
	var cod string
	fmt.Printf("codigo: ")
	fmt.Scanf("%s\n", &cod)

	apresentacoes := controladoras.RecuperarApresentacoes(cod)
	e := controladoras.RecuperarEvento(cod)

	utils.ClearScreen()

	fmt.Printf("\tVisualizando Apresentações do Evento: %s\n\n", e.Nome)
	for _, apr := range apresentacoes {
		fmt.Println(strings.Repeat("-", 10))
		fmt.Println(apr.ToString())
	}
	fmt.Println()

	utils.Pause()
}
