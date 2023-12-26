package config

import (
	"database/sql"
	"fmt"
	"os"
)

func GetDSN() string {
	dbUser := os.Getenv("DB_USER")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	return fmt.Sprintf("%s@tcp(%s:%s)/%s", dbUser, dbHost, dbPort, dbName)
}

func CreateDB(db *sql.DB) {
	// _, err := db.Query("CREATE DATABASE test")
	// if err != nil {
	// 	panic(err.Error())
	// }

	_, err := db.Query("CREATE TABLE company (id integer AUTO_INCREMENT UNIQUE, name varchar(255) NOT NULL, PRIMARY KEY(id))")
	if err != nil {
		panic(err.Error())
	}

	_, err = db.Query("CREATE TABLE product (id integer AUTO_INCREMENT UNIQUE, code varchar(255) NOT NULL, name varchar(255) NOT NULL, brand varchar(255), detail varchar(255), company_id integer NOT NULL, PRIMARY KEY(id), CONSTRAINT fk_company FOREIGN KEY(company_id) REFERENCES `company`(`id`))")
	if err != nil {
		panic(err.Error())
	}

	_, err = db.Query("CREATE TABLE movement (id integer AUTO_INCREMENT UNIQUE, date varchar(50) NOT NULL, shipping_code varchar(255) NOT NULL, pallets integer, units integer, deposito varchar(255), PRIMARY KEY(id))")
	if err != nil {
		panic(err.Error())
	}

	_, err = db.Query("CREATE TABLE movements_products (id integer AUTO_INCREMENT UNIQUE, movement_id integer NOT NULL, product_id integer NOT NULL, PRIMARY KEY(id), CONSTRAINT fk_movement FOREIGN KEY(movement_id) REFERENCES `movement`(`id`), CONSTRAINT fk_product FOREIGN KEY(product_id) REFERENCES `product`(`id`))")
	if err != nil {
		panic(err.Error())
	}
}
