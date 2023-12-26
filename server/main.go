package main

import (
	"database/sql"
	"fmt"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"github.com/stock-controller/app/config"
	"github.com/stock-controller/app/router"
)

func main() {
	app := gin.Default()

	router.GetRouter(app)

	if err := godotenv.Load(".env"); err != nil {
		panic(err.Error())
	}

	db, err := sql.Open("mysql", config.GetDSN())
	if err != nil {
		panic(err.Error())
	}

	if err = db.Ping(); err != nil {
		config.CreateDB(db)
	} else {
		fmt.Println("Conexion a db exitosa")
		config.CreateDB(db)
	}
	defer db.Close()

	app.Run("localhost:8080")
}
