package book

import (
	"fmt"

	"github.com/charathram/fiber101/database"
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"github.com/jinzhu/gorm"
)

// Book represents a book
type Book struct {
	gorm.Model
	Title  string `json:"title"`
	Author string `json:"author"`
	Rating int    `json:"rating"`
}

// GetBooks returns a list of all books
func GetBooks(c *fiber.Ctx) error {
	db := database.DBConn
	var books []Book

	db.Find(&books)
	return c.JSON(books)
}

// GetBook returns details of a single book
func GetBook(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DBConn
	var book Book

	db.Find(&book, id)

	return c.JSON(book)
}

// NewBook creates a new book in the DB
func NewBook(c *fiber.Ctx) error {

	var t interface{}

	db := database.DBConn

	t = c.Locals("user")
	fmt.Printf("%T\n", t)

	switch t := t.(type) {
	default:
		fmt.Println("Default")
	case *jwt.Token:
		fmt.Println("JWT Token Pointer")
		myToken := *t
		fmt.Println(myToken.Claims)
	}

	book := new(Book)
	if err := c.BodyParser(book); err != nil {
		return c.Status(fiber.ErrBadRequest.Code).JSON(err.Error())
	}

	db.Create(&book)
	return c.JSON(book)
}

// DeleteBook deletes a book from the DB
func DeleteBook(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DBConn
	var book Book

	db.First(&book, id)
	if book.Title == "" {
		return c.Status(fiber.ErrNotFound.Code).JSON(fiber.Map{"error": "No book found with given ID"})
	}
	db.Delete(&book)
	return c.SendString("Book successfully deleted")
}
