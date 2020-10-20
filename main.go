package main

import (
	"fmt"

	"github.com/charathram/fiber101/auth"
	"github.com/charathram/fiber101/book"
	"github.com/charathram/fiber101/database"
	"github.com/charathram/fiber101/middleware"
	"github.com/gofiber/fiber/v2"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func initDatabase() {
	var err error
	database.DBConn, err = gorm.Open("sqlite3", "books.db")
	if err != nil {
		panic("Unable to connect to database")
	}
	fmt.Println("Database connection successfully opened")
	database.DBConn.AutoMigrate(&book.Book{})
	fmt.Println("Database migrated")
}

func setupRoutes(app *fiber.App) {
	app.Post("/login", auth.Login)
	app.Get("/api/v1/book", book.GetBooks)
	app.Get("/api/v1/book/:id", book.GetBook)
	app.Post("/api/v1/book", middleware.Protected(), book.NewBook)
	app.Delete("/api/v1/book/:id", middleware.Protected(), book.DeleteBook)
}

func main() {
	app := fiber.New()
	initDatabase()
	defer database.DBConn.Close()

	setupRoutes(app)

	app.Listen(":3000")
}
