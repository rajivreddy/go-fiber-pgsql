package handlers

import (
	"log"

	fiber "github.com/gofiber/fiber/v2"
	"github.com/rajivreddy/go-fiber-pgsql/pkg/datatypes"
	"github.com/rajivreddy/go-fiber-pgsql/pkg/postgres"
	"golang.org/x/crypto/bcrypt"
)

func Hash(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func Verify(hashed, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(password))
	return err == nil
}

func GetUser(c *fiber.Ctx) error {
	username := c.Params("username")
	log.Printf("Username: %v", username)
	user, err := postgres.GetUser(username)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.Status(200).JSON(fiber.Map{
		"user": user,
	})

}

func CreateUser(c *fiber.Ctx) error {
	users := datatypes.User{}

	if err := c.BodyParser(&users); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	log.Println(users)
	var err error
	users.Password, err = Hash(users.Password)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	id, err := postgres.CreateUser(users)
	if id == 0 || err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.Status(201).JSON(fiber.Map{
		"message": "User has been created",
		"id":      id,
	})
}
