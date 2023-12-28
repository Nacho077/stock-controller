package config

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func GetDB() *sql.DB {
	dbName := os.Getenv("DB_NAME")

	db, err := sql.Open("mysql", getDSN(dbName))
	if err != nil {
		panic(err.Error())
	}

	createDB(db, dbName)

	if err = db.Ping(); err != nil {
		panic(err.Error())
	}

	return db
}

func getDSN(dbName string) string {
	dbUser := os.Getenv("DB_USER")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")

	fmt.Println(dbUser)

	return fmt.Sprintf("%s@tcp(%s:%s)/%s", dbUser, dbHost, dbPort, dbName)
}

func createDB(db *sql.DB, dbName string) {

	_, err := db.Query(fmt.Sprintf("CREATE TABLE IF NOT EXISTS %s.company (id integer AUTO_INCREMENT UNIQUE, name varchar(255) NOT NULL UNIQUE, PRIMARY KEY(id))", dbName))
	if err != nil {
		panic(err.Error())
	}

	_, err = db.Query(fmt.Sprintf("CREATE TABLE IF NOT EXISTS %s.product (id integer AUTO_INCREMENT UNIQUE, code varchar(255) NOT NULL, name varchar(255) NOT NULL, brand varchar(255), detail varchar(255), company_id integer NOT NULL, PRIMARY KEY(id), CONSTRAINT fk_company FOREIGN KEY(company_id) REFERENCES %s.company(id))", dbName, dbName))
	if err != nil {
		panic(err.Error())
	}

	_, err = db.Query(fmt.Sprintf("CREATE TABLE IF NOT EXISTS %s.movement (id integer AUTO_INCREMENT UNIQUE, date varchar(50) NOT NULL, shipping_code varchar(255) NOT NULL, pallets integer, units integer, deposit varchar(255), observations varchar(255), PRIMARY KEY(id))", dbName))
	if err != nil {
		panic(err.Error())
	}

	_, err = db.Query(fmt.Sprintf("CREATE TABLE IF NOT EXISTS %s.movements_products (id integer AUTO_INCREMENT UNIQUE, movement_id integer NOT NULL, product_id integer NOT NULL, PRIMARY KEY(id), CONSTRAINT fk_movement FOREIGN KEY(movement_id) REFERENCES %s.movement(id), CONSTRAINT fk_product FOREIGN KEY(product_id) REFERENCES %s.product(id))", dbName, dbName, dbName))
	if err != nil {
		panic(err.Error())
	}
}
