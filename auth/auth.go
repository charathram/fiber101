package auth

import (
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
)

// Login logs the user in and returns a token if successfully
// authenticated
func Login(c *fiber.Ctx) error {
	type User struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	var user User

	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.ErrBadRequest.Code).JSON(fiber.Map{"status": "error", "message": "Error on login request", "data": err})
	}
	if user.Email != "charath@example.com" || user.Password != "abcd1234" {
		return c.Status(fiber.ErrUnauthorized.Code).JSON(fiber.Map{"status": "error", "message": "Invalid email or password", "data": nil})
	}

	// if the user is authorized, create a new token
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	//claims["email"] = user.Email
	claims["email"] = "test"
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	t, err := token.SignedString([]byte("goisfun"))
	if err != nil {
		return c.Status(fiber.ErrInternalServerError.Code).JSON(fiber.Map{"status": "error", "message": "Internal server error", "data": nil})
	}
	return c.JSON(fiber.Map{"status": "success", "message": "Successful login", "data": t})
}
