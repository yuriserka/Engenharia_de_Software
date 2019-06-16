package utils

import (
	"bufio"
	"fmt"
	"os"
	"reflect"

	"github.com/howeyc/gopass"
)

// GetUserData recebe o input do usuario tanto no login quanto no cadastro
func GetUserData() (cpf string, senha []byte) {
	fmt.Printf("Cpf: ")
	fmt.Scanf("%s\n", &cpf)

	fmt.Printf("Senha: ")
	senha, err := gopass.GetPasswdMasked()

	if err != nil {
		panic(err.Error())
	}

	return
}

// getStringInput remove o \n quando esta pegando input do usuario e retorna esta string
func getStringInput(msg string) string {
	reader := bufio.NewReader(os.Stdin)

	fmt.Printf("%s: ", msg)
	s, err := reader.ReadString('\n')
	if err != nil {
		return ""
	}
	s = s[:len(s)-2]
	return s
}

// GetNStringInputs recebe como argumento uma struct que contem as variaveis string e qual
// a mensagem que deve aparecer para elas, e também a qtd de variaveis
func GetNStringInputs(i interface{}, inputMsgs []string, numVars int) interface{} {
	vptr := reflect.Indirect(reflect.ValueOf(i))
	for i := 0; i < numVars; i++ {
		si := getStringInput(inputMsgs[i])
		vptr.Field(i).SetString(si)
	}

	return vptr.Interface()
}
