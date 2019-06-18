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

func visualizarTodosEventos() error {
	utils.ClearScreen()
	fmt.Printf("\tVisualizando Todos so Eventos\n\n")

	eventos, e1 := controladoras.RecuperarEventos()
	if e1 != nil {
		fmt.Println(e1.Error())
		return e1
	}
	for _, e := range eventos {
		fmt.Println(strings.Repeat("-", 10))
		fmt.Println(e.ToString())
	}
	fmt.Println()
	return nil
}

func cadastrarApresentacao() {
	if noEvents := visualizarTodosEventos(); noEvents != nil {
		utils.Pause()
		return
	}

	fmt.Printf("\tEscolha um Evento pelo Código\n\n")
	var cod string
	fmt.Printf("codigo: ")
	fmt.Scanf("%s\n", &cod)

	_, err := controladoras.RecuperarEvento(cod)
	if err != nil {
		fmt.Println(err.Error())
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
		a, e1 := controladoras.RecuperarApresentacao(v.Codigo)
		if e1 != nil {
			fmt.Println(e1.Error())
			return
		}
		if err := controladoras.LigarApresentacaoEvento(cod, a); err != nil {
			fmt.Println(err.Error())
			return
		}
		fmt.Println("Apresentação cadastrada com sucesso")
	}

	utils.Pause()
}

func visualizarApresentacoes() {
	if noEvents := visualizarTodosEventos(); noEvents != nil {
		utils.Pause()
		return
	}

	fmt.Printf("\tEscolha um Evento pelo Código\n\n")
	var cod string
	fmt.Printf("codigo: ")
	fmt.Scanf("%s\n", &cod)

	apresentacoes, e1 := controladoras.RecuperarApresentacoes(cod)
	if e1 != nil {
		fmt.Println(e1.Error())
		utils.Pause()
		return
	}
	ev, e2 := controladoras.RecuperarEvento(cod)
	if e2 != nil {
		fmt.Println(e2.Error())
		utils.Pause()
		return
	}

	utils.ClearScreen()

	fmt.Printf("\tVisualizando Apresentações do Evento: %s\n\n", ev.Nome)
	for _, apr := range apresentacoes {
		fmt.Println(strings.Repeat("-", 10))
		fmt.Println(apr.ToString())
	}
	fmt.Println()

	utils.Pause()
}
