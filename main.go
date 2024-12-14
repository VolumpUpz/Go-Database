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
	databaseName = "mydatabase"
	username     = "myuser"
	password     = "mypassword"
)

var db *sql.DB

type Product struct {
	ID    int
	Name  string
	Price int
}

type ProductWithSupplier struct {
	ID           int
	Name         string
	Price        int
	SupplierName string
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

	// err = createProduct(&Product{
	// 	Name:  "Go Product",
	// 	Price: 222,
	// })

	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Println("Create product successful")

	product, err := getProduct(1)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(product)

	//err = updateProduct(3, &Product{Name: "Go Products", Price: 224})
	// product, err = updateProductWithRetuning(3, &Product{Name: "Go Productss", Price: 225})
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(product)

	// err = deleteProduct(3)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	products, err := getProducts()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(products)

	ps, err := getProductWithSupplier()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(ps)

}

func createProduct(product *Product) error {

	_, err := db.Exec("INSERT INTO PRODUCTS (name,price) VALUES($1,$2);", product.Name, product.Price)
	return err
}

func getProduct(id int) (Product, error) {
	var p Product
	row := db.QueryRow("SELECT id,name,price FROM PRODUCTS WHERE id = $1;", id)

	err := row.Scan(&p.ID, &p.Name, &p.Price)
	if err != nil {
		return Product{}, err
	}
	return p, nil

}

func updateProduct(id int, product *Product) error {
	_, err := db.Exec("UPDATE PRODUCTS SET name=$1,price=$2 WHERE id = $3;", product.Name, product.Price, id)
	return err
}

func updateProductWithRetuning(id int, product *Product) (Product, error) {
	var p Product

	row := db.QueryRow("UPDATE PRODUCTS SET name=$1,price=$2 WHERE id = $3 RETURNING id, name, price;", product.Name, product.Price, id)
	err := row.Scan(&p.ID, &p.Name, &p.Price)
	if err != nil {
		return Product{}, err
	}
	return p, nil

}

func deleteProduct(id int) error {
	_, err := db.Exec("DELETE FROM PRODUCTS WHERE id= $1;", id)
	return err
}

func getProducts() ([]Product, error) {

	rows, err := db.Query("SELECT id,name,price from PRODUCTS")
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var products []Product //slice
	for rows.Next() {      //เป็นการ shift iteration ของ rows ไปเรื่อยๆ จนจบ
		var p Product
		err := rows.Scan(&p.ID, &p.Name, &p.Price)
		if err != nil {
			return nil, err
		}
		products = append(products, p)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return products, nil
}

func getProductWithSupplier() ([]ProductWithSupplier, error) {

	rows, err := db.Query(`SELECT p.id, p.name, p.Price, s.Name as SupplierName
						   FROM PRODUCTS p 
						   INNER JOIN SUPPLIER s 
						   ON p.supplier_id = s.id
						   WHERE 1 = 1
						 `)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var productWithSuppliers []ProductWithSupplier

	for rows.Next() {
		var ps ProductWithSupplier
		err := rows.Scan(&ps.ID, &ps.Name, &ps.Price, &ps.SupplierName)
		if err != nil {
			return nil, err
		}
		productWithSuppliers = append(productWithSuppliers, ps)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}
	return productWithSuppliers, nil

}
