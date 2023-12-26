package config

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func GetDB() *sql.DB {
	db, err := sql.Open("mysql", getDSN())
	if err != nil {
		panic(err.Error())
	}

	createDB(db)

	if err = db.Ping(); err != nil {
		panic(err.Error())
	}

	return db
}

func getDSN() string {
	dbUser := os.Getenv("DB_USER")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	return fmt.Sprintf("%s@tcp(%s:%s)/%s", dbUser, dbHost, dbPort, dbName)
}

func createDB(db *sql.DB) {
	_, err := db.Exec("CREATE DATABASE IF NOT EXISTS stock_controller")
	if err != nil {
		panic(err.Error())
	}

	_, err = db.Exec("USE stock_controller")

	_, err = db.Query("CREATE TABLE IF NOT EXISTS company (id integer AUTO_INCREMENT UNIQUE, name varchar(255) NOT NULL, PRIMARY KEY(id))")
	if err != nil {
		panic(err.Error())
	}

	_, err = db.Query("CREATE TABLE IF NOT EXISTS product (id integer AUTO_INCREMENT UNIQUE, code varchar(255) NOT NULL, name varchar(255) NOT NULL, brand varchar(255), detail varchar(255), company_id integer NOT NULL, PRIMARY KEY(id), CONSTRAINT fk_company FOREIGN KEY(company_id) REFERENCES `company`(`id`))")
	if err != nil {
		panic(err.Error())
	}

	_, err = db.Query("CREATE TABLE IF NOT EXISTS movement (id integer AUTO_INCREMENT UNIQUE, date varchar(50) NOT NULL, shipping_code varchar(255) NOT NULL, pallets integer, units integer, deposito varchar(255), PRIMARY KEY(id))")
	if err != nil {
		panic(err.Error())
	}

	_, err = db.Query("CREATE TABLE IF NOT EXISTS movements_products (id integer AUTO_INCREMENT UNIQUE, movement_id integer NOT NULL, product_id integer NOT NULL, PRIMARY KEY(id), CONSTRAINT fk_movement FOREIGN KEY(movement_id) REFERENCES `movement`(`id`), CONSTRAINT fk_product FOREIGN KEY(product_id) REFERENCES `product`(`id`))")
	if err != nil {
		panic(err.Error())
	}
}
