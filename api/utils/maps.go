package utils

import (
	"fmt"
	"sort"

	"github.com/yuriserka/Engenharia_de_Software/api/entidades"
)

// OrdenaMap ordena os indices do map para que possa imprimir na ordem certa, é util
// para os menus
func OrdenaMap(m map[int]string) []int {
	sortedIndexes := make([]int, 0, len(m))
	for i := range m {
		sortedIndexes = append(sortedIndexes, i)
	}
	sort.Ints(sortedIndexes)

	return sortedIndexes
}

func AddCartaoCredito(m map[string]map[string]*entidades.CartaoDeCredito, key1,
	key2 string, obj *entidades.CartaoDeCredito) {
	mm, existe := m[key1]
	if !existe {
		mm = make(map[string]*entidades.CartaoDeCredito)
		m[key1] = mm
	}
	_, ok := mm[key2]
	if !ok {
		mm[key2] = obj
	} else {
		fmt.Println("Cartao ja cadastrado")
	}
}
