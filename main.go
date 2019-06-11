package main

import (
	servicos "github.com/yuriserka/Engenharia_de_Software/backend/controladoras_servicos"
)

func main() {
	ctrlSControle := &servicos.ControladoraServicoControle{}
	ctrlSControle.Build()
}
