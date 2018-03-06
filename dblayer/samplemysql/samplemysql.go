package main

import (
	"database/sql"
	"fmt"

	"log"

	_ "github.com/go-sql-driver/mysql"
)

type animal struct {
	id         int
	animalType string
	nickname   string
	zona       int
	age        int
}

func main() {
	log.Println("Connecting ...database ")
	db, err := sql.Open("mysql", "gouser:!password@/go_portal")
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	log.Println("Querying with arguments")

	rows, err := db.Query("SELECT * FROM animal WHERE age>?", 1)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	animals := []animal{}
	for rows.Next() {
		a := animal{}
		err := rows.Scan(&a.id, &a.animalType, &a.nickname, &a.zona, &a.age)
		if err != nil {
			log.Fatal(err)
		}
		animals = append(animals, a)
	}
	fmt.Println(animals)
}
