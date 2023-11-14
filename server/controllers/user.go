package controllers

import (
	"api-estadia-express/init/db"
	"api-estadia-express/init/logger"
	"api-estadia-express/internal"
	"api-estadia-express/models"
	"encoding/json"
	"github.com/gofiber/fiber/v2"
)

func CreateUser(c *fiber.Ctx) error {
	var user models.User

	if err := json.Unmarshal(c.Body(), &user); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("invalid user data")
	}

	hashPassword, err := internal.HashString(user.Password)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("error to hash user password")
	}

	user.Password = hashPassword

	if err = user.CreateUser(db.InstanceDB()); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("error to create user")
	}

	return c.Status(fiber.StatusOK).JSON(user)
}

func Login(c *fiber.Ctx) error {
	var payload models.LoginPayload
	var user models.User

	if err := json.Unmarshal(c.Body(), &payload); err != nil {
		logger.Error("invalid fields", err)
		return c.Status(fiber.StatusBadRequest).SendString("invalid fields")
	}

	user.Email = payload.Email
	userList, err := user.FindUsersWithFilter(db.InstanceDB())

	if err != nil || len(userList) == 0 {
		logger.Error("user not found", err)
		return c.Status(fiber.StatusBadRequest).SendString("user not found")
	}

	if err = userList[0].CheckPassword(payload.Password); err != nil {
		logger.Error("invalid credentials", err)
		return c.Status(fiber.StatusBadRequest).SendString("invalid credentials")
	}

	jwtWrapper := internal.JwtWrapper{
		SecretKey:       "verysecretkey",
		Issuer:          "AuthService",
		ExpirationHours: 24,
	}

	signedToken, err := jwtWrapper.GenerateToken(user.Email)

	if err != nil {
		logger.Error("error to create jwt token", err)
		return c.Status(fiber.StatusBadRequest).SendString("error to create token")
	}

	tokenResponse := models.LoginResponse{
		Token: signedToken,
	}

	return c.Status(fiber.StatusOK).JSON(tokenResponse)
}
