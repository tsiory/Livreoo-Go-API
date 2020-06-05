package product

import (
	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
	"github.com/tsiory/livreoo-go-api/database"
)

// Product Model
type Product struct {
	gorm.Model
	Label string `json:"label"`
	Price string `json:"price"`
}

// Get All products from the database
func GetProducts(c *fiber.Ctx) {
	db := database.DBConn
	var products []Product
	db.Find(&products)
	c.JSON(products)
}

// Get one product with :id from the database
func GetProduct(c *fiber.Ctx) {
	id := c.Params("id")
	db := database.DBConn
	var product Product
	db.Find(&product, id)
	c.JSON(product)
}

// Add new product into the database
func NewProduct(c *fiber.Ctx) {
	db := database.DBConn
	var product Product
	if err := c.BodyParser(&product); err != nil {
		c.Status(503).Send(err)
		return
	}
	db.Create(&product)
	c.JSON(product)
}

// Delete the product with the given :id from database
func DeleteProduct(c *fiber.Ctx) {
	id := c.Params("id")
	db := database.DBConn
	var product Product
	db.Find(&product, id)
	if product.Label == "" {
		c.Status(404).SendString("Product Not Found.")
		return
	}
	db.Delete(&product, id)
	c.JSON(&product)
}
