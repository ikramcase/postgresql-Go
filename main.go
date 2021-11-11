package main

import (
	"database/sql"
	"encoding/json"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "root"
	dbname   = "ik_demo"
)

type Student struct {
	ID   int    `json:"id,omitemp"`
	Nama string `json:"nama,omitemp"`
}

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")

	sql := "select id,name FROM name "
	data, err := DB.Query(sql)
	if err != nil {
		saveError := fmt.Sprintf("Error Query, and %s", err)
	}
	type ManyStudents []Student

	var manyStudents ManyStudents
	for data.Next() {
		var perStudent Student
		err = data.Scan(&perStudent.ID, &perStudent.Nama)
		if err != nil {
			saveError := fmt.Sprintf("Error Looping data, and %s", err)
		}
		manyStudents = append(manyStudents, perStudent)
	}
	err := json.NewDecoder(r.Body).Decode(ManyStudents)
	if err != nil {
		util.Respond(w, util.MetaMsg(false, "Invalid request"))
	}
}
