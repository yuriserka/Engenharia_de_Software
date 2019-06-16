package entidades

import (
	"fmt"
)

// CartaoDeCredito é a classe que servirá para validar se o usuário pode ou não pagar pela apresentação.
type CartaoDeCredito struct {
	Numero, Codigo, Validade string
}

// ToString facilita a impressão desta entidade
func (c CartaoDeCredito) ToString() string {
	return fmt.Sprintf("Numero: %s\nCódigo %s\nValidade: %s\n", c.Numero, c.Codigo, c.Validade)
}
