package main

import (
	servicos "github.com/Engenharia_de_Software/backend/controladoras_servicos"
)

/* Executar para funcionar =>
 * $ go mod download
 * $ go run .
 */
func main() {
	ctrlSControle := &servicos.ControladoraServicoControle{}
	ctrlSControle.Build()
}
