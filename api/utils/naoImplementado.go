package utils

import "fmt"

// NaoImplementado serve para deixar de forma clara o que ainda não foi feito dentro do sistema
func NaoImplementado(f string) {
	fmt.Println("\tFunção", f, "ainda não foi implementada...")
	Pause()
}
