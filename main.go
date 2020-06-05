package main

import (
	"fmt"

	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/tsiory/livreoo-go-api/database"
	"github.com/tsiory/livreoo-go-api/product"
)

func main() {
	app := fiber.New()
	initDatabase()
	defer database.DBConn.Close()

	setupRoutes(app)

	app.Listen(3000)
}

func setupRoutes(app *fiber.App) {
	app.Get("/api/v1/products", product.GetProducts)
	app.Get("/api/v1/products/:id", product.GetProduct)
	app.Post("/api/v1/products", product.NewProduct)
	app.Delete("/api/v1/products/:id", product.DeleteProduct)
}

func initDatabase() {
	var err error
	database.DBConn, err = gorm.Open("sqlite3", "products.db")
	if err != nil {
		panic(err)
	}
	fmt.Println("Database connection successfully opened.")
	database.DBConn.AutoMigrate(&product.Product{})
	fmt.Println("Database Migrated.")
}
