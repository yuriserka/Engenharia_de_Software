package main

import "fmt"

/* Executar para funcionar =>
 * $ go get ./...
 * $ go run .
 */
func main() {
	fmt.Println(func() string {
		msg := func() string {
			return "Vamos fazer tudo "
		}() + func() string {
			return "Funcional"
		}()
		return msg
	}(), "UHU")
}
