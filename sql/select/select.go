package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type usuario struct {
	id   int
	nome string
}

func main() {
	db, err := sql.Open("mysql", "root:Fsg@250583@/cursogo")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	rows, _ := db.Query("select id , nome from usuarios where id > ?", 0)
	defer rows.Close()

	usuarioSlice := make([]usuario, 10)
	contador := 0
	for rows.Next() {
		var u usuario
		rows.Scan(&u.id, &u.nome)
		fmt.Println(u)
		usuarioSlice[contador] = u
		contador++
	}

	fmt.Println(usuarioSlice[0].id)
	fmt.Println(usuarioSlice[0].nome)
}
