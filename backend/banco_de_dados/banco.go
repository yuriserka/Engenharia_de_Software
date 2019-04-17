package backend

import (
	"database/sql"
	"fmt"

	// ta vazio pq nao importa o nome que vai ser importado
	_ "github.com/go-sql-driver/mysql"
)

type bancoDeDados struct {
	db *sql.DB
}

// Init inicia a conexão com o banco de dados. Melhorar isso ainda.
func (d *bancoDeDados) Init() {
	password := "NONE"
	bdName := "NONE"
	db, err := sql.Open("mysql", "root:"+password+"@tcp(127.0.0.1)/"+bdName)

	d.db = db

	if err != nil {
		panic(err.Error())
	}

	fmt.Println("Conectado com sucesso no banco de dados")
}

// Close fecha a conexão com o banco de dados.
func (d *bancoDeDados) Close() {
	defer d.db.Close()

	fmt.Println("banco de dados fechado com sucesso")
}
