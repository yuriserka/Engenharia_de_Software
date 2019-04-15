package utils

import (
	"os"
	"os/exec"
	"runtime"
)

var clear map[string]func() //create a map for storing clear funcs

func init() {
	clear = make(map[string]func()) //Initialize it
	clear["linux"] = func() {
		cmd := exec.Command("clear") //Linux example, its tested
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
	clear["windows"] = func() {
		cmd := exec.Command("cmd", "/c", "cls") //Windows example, its tested
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}

// ClearScreen limpa a tela do terminal.
func ClearScreen() {
	if executeClean, supportedPlatform := clear[runtime.GOOS]; supportedPlatform {
		executeClean()
	} else {
		panic("Plataforma não suportada, impossível limpar tela.")
	}
}
