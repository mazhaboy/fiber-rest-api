package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)
func main(){
	connStr:="user=postgres password=Knopka1 dbname=postgres sslmode=disable"
	db, err:=sql.Open("postgres", connStr)
	if err!=nil{
		log.Fatal(err)
	}
	defer db.Close()
	fmt.Println("Successfully connected to the database")
	result , err := db.Exec("insert into Products (model, company, price) values ('iPhone', $1,$2)", "Apple", 72000)
	if err!=nil{
		log.Fatal(err)
	}
	fmt.Println("Successfully inserted to the database")
	fmt.Println(result.LastInsertId())
}