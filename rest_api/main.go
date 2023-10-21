package main

import (
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type User struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Email   string `json:"email"`
	Address string `json:"address"`
}

type Response struct {
	Success    bool   `json:"success"`
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
}

type GetResponse struct {
	Success    bool   `json:"success"`
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
	Payload    []User `json:"payload"`
}

var users = []User{
	{
		ID:      1,
		Name:    "Muhammad Ikhsan Hilmi",
		Email:   "ikhsanhilmimuhammad@gmail.com",
		Address: "Bandung",
	},
}

func getLogMessage(message string, c *fiber.Ctx, traceID string) string {
	currentTime := time.Now().Format("2006/01/02 15:04:05")
	return currentTime + " message=\"" + message + "\" method=" + c.Method() + " uri=" + c.Path() + " trace_id=" + traceID
}

func Trace() fiber.Handler {
	return func(c *fiber.Ctx) error {
		traceID := uuid.New().String()
		c.Locals("trace_id", traceID)

		log.Print(getLogMessage("incoming request", c, traceID))

		c.Set("X-Request-ID", traceID)
		return c.Next()
	}
}

func createUser(c *fiber.Ctx) error {
	traceID := c.Locals("trace_id").(string)

	var newUser User
	if err := c.BodyParser(&newUser); err != nil {
		log.Print(getLogMessage("error when trying to create user", c, traceID))
		return err
	}

	newUser.ID = len(users) + 1
	users = append(users, newUser)

	log.Print(getLogMessage("finish request", c, traceID))

	return c.Status(http.StatusCreated).JSON(Response{
		Success:    true,
		StatusCode: http.StatusCreated,
		Message:    "created success",
	})
}

func getAllUser(c *fiber.Ctx) error {
	traceID := c.Locals("trace_id").(string)

	log.Print(getLogMessage("finish request", c, traceID))

	return c.Status(http.StatusOK).JSON(GetResponse{
		Success:    true,
		StatusCode: http.StatusOK,
		Message:    "get all success",
		Payload:    users,
	})
}

func updateUser(c *fiber.Ctx) error {
	traceID := c.Locals("trace_id").(string)

	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		log.Print(getLogMessage("error when trying to update user", c, traceID))
		return err
	}

	var updatedUser User
	if err := c.BodyParser(&updatedUser); err != nil {
		log.Print(getLogMessage("error when trying to update user", c, traceID))
		return err
	}

	var userFound bool
	for i, user := range users {
		if user.ID == id {
			users[i] = updatedUser
			userFound = true
			break
		}
	}

	if !userFound {
		log.Print(getLogMessage("error when try to get users with no data", c, traceID))
		return c.Status(http.StatusNotFound).JSON(Response{
			Success:    false,
			StatusCode: http.StatusNotFound,
			Message:    "user not found",
		})
	}

	log.Print(getLogMessage("finish request", c, traceID))

	return c.Status(http.StatusOK).JSON(Response{
		Success:    true,
		StatusCode: http.StatusOK,
		Message:    "update success",
	})
}

func deleteUser(c *fiber.Ctx) error {
	traceID := c.Locals("trace_id").(string)

	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		log.Print(getLogMessage("error when trying to delete user", c, traceID))
		return err
	}

	var userFound bool
	for i, user := range users {
		if user.ID == id {
			users = append(users[:i], users[i+1:]...)
			userFound = true
			break
		}
	}

	if !userFound {
		log.Print(getLogMessage("error when try to get users with no data", c, traceID))
		return c.Status(http.StatusNotFound).JSON(Response{
			Success:    false,
			StatusCode: http.StatusNotFound,
			Message:    "user not found",
		})
	}

	log.Print(getLogMessage("finish request", c, traceID))

	return c.Status(http.StatusOK).JSON(Response{
		Success:    true,
		StatusCode: http.StatusOK,
		Message:    "delete success",
	})
}

func main() {
	app := fiber.New()

	app.Use(Trace())

	v1 := app.Group("/api/v1")

	userRouter := v1.Group("/users")
	userRouter.Post("/", createUser)
	userRouter.Get("/", getAllUser)
	userRouter.Put("/:id", updateUser)
	userRouter.Delete("/:id", deleteUser)

	log.Fatal(app.Listen(":4444"))
}
