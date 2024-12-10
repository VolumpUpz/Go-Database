package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

const (
	host         = "localhost"
	port         = 5432
	databaseName = "my-db"
	username     = "myuser"
	password     = "mypassword"
)

var db *sql.DB

type Product struct {
	ID    int
	Name  string
	Price int
}

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, username, password, databaseName)

	sdb, err := sql.Open("postgres", psqlInfo)

	if err != nil {
		log.Fatal("err 1 :", err)
	}
	db = sdb

	err = db.Ping()

	if err != nil {
		log.Fatal("err 2 :", err)
	}

	fmt.Println("Connection Database Successful")

	err = createProduct(&Product{
		Name:  "Go Product",
		Price: 222,
	})

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Create product successful")

}

func createProduct(product *Product) error {

	_, err := db.Exec("INSERT INTO PRODUCTS (name,price) VALUES($1,$2);", product.Name, product.Price)
	return err
}
