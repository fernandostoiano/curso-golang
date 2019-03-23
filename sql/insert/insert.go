package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:Fsg@250583@/cursogo")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	stmt, err := db.Prepare("insert into usuarios(nome) values(?)")
	if err != nil {
		panic(err)
	}
	stmt.Exec("Maria")
	stmt.Exec("João")

	res, _ := stmt.Exec("Pedro")

	id, _ := res.LastInsertId()
	fmt.Println(id)

	linhas, _ := res.RowsAffected()
	fmt.Println(linhas)
}
