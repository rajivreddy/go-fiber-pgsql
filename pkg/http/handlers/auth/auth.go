package handlers

import (
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	users "github.com/rajivreddy/go-fiber-pgsql/pkg/http/handlers/users"
	"github.com/rajivreddy/go-fiber-pgsql/pkg/postgres"
)

func Login(c *fiber.Ctx) error {
	type login struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	var loginInput login
	if err := c.BodyParser(&loginInput); err != nil {
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	log.Printf("Username: %s, Password: %s", loginInput.Username, loginInput.Password)

	// Check if username and password are correct

	user, error := postgres.GetUser(loginInput.Username)
	if error != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": error.Error(),
		})
	}
	if !users.Verify(user.Password, loginInput.Password) {
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	log.Printf("User: %v", user)

	// Generate JWT token
	token := jwt.New(jwt.SigningMethodHS256)
	log.Printf("Token: %v", token)

	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = loginInput.Username
	claims["admin"] = true
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	log.Printf("Claims: %v", claims)

	toeknData, err := token.SignedString([]byte("secret"))
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "Login successful",
		"token":   toeknData,
	})
}
