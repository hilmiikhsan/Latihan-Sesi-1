package product

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
)

type ResponseSuccess struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

type ResponseWithPayloadSuccess struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Payload interface{} `json:"payload"`
}

type ResponseError struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Error   string `json:"error"`
}

type Handler struct {
	svc Service
}

func NewHandler(svc Service) Handler {
	return Handler{
		svc: svc,
	}
}

func (h Handler) CreateProduct(c *fiber.Ctx) error {
	var req CreateOrUpdateProductRequest

	err := c.BodyParser(&req)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(ResponseError{
			Success: false,
			Message: "ERR BAD REQUEST",
			Error:   err.Error(),
		})
	}

	model := Product{
		Name:     req.Name,
		Category: req.Category,
		Price:    req.Price,
		Stock:    req.Stock,
	}

	err = h.svc.CreateProduct(c.UserContext(), model)
	if err != nil {
		var payload ResponseError
		httpCode := 400
		switch err {
		case ErrEmptyName, ErrEmptyCategory, ErrEmptyPrice, ErrEmptyStock:
			payload = ResponseError{
				Success: false,
				Message: "ERR BAD REQUEST",
				Error:   err.Error(),
			}
			httpCode = http.StatusBadRequest
		default:
			payload = ResponseError{
				Success: false,
				Message: "ERR INTERNAL",
				Error:   err.Error(),
			}
			httpCode = http.StatusInternalServerError
		}
		return c.Status(httpCode).JSON(payload)
	}
	return c.Status(fiber.StatusCreated).JSON(ResponseSuccess{
		Success: true,
		Message: "CREATE SUCCESS",
	})
}

func (h Handler) GetAllProduct(c *fiber.Ctx) error {
	res, err := h.svc.GetAllProduct(c.UserContext())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(ResponseError{
			Success: false,
			Message: "ERR INTERNAL",
			Error:   err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(ResponseWithPayloadSuccess{
		Success: true,
		Message: "GET ALL SUCCESS",
		Payload: res,
	})
}

func (h Handler) GetProductById(c *fiber.Ctx) error {
	idStr := c.Params("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(ResponseError{
			Success: false,
			Message: "ERR BAD REQUEST",
			Error:   "invalid id",
		})

	}

	res, err := h.svc.GetProductById(c.UserContext(), id)
	if err != nil {
		if strings.Contains(err.Error(), "record not found") {
			return c.Status(fiber.StatusNotFound).JSON(ResponseError{
				Success: false,
				Message: "ERR NOT FOUND",
				Error:   "data tidak ditemukan",
			})
		}

		return c.Status(fiber.StatusInternalServerError).JSON(ResponseError{
			Success: false,
			Message: "ERR INTERNAL",
			Error:   err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(ResponseWithPayloadSuccess{
		Success: true,
		Message: "GET DATA SUCCESS",
		Payload: res,
	})
}

func (h Handler) UpdateProduct(c *fiber.Ctx) error {
	idStr := c.Params("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(ResponseError{
			Success: false,
			Message: "ERR BAD REQUEST",
			Error:   "invalid id",
		})

	}

	var req CreateOrUpdateProductRequest

	err = c.BodyParser(&req)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(ResponseError{
			Success: false,
			Message: "ERR BAD REQUEST",
			Error:   err.Error(),
		})
	}

	model := Product{
		Name:     req.Name,
		Category: req.Category,
		Price:    req.Price,
		Stock:    req.Stock,
	}

	err = h.svc.UpdateProduct(c.UserContext(), id, model)
	if err != nil {
		var payload ResponseError
		httpCode := 400
		switch {
		case err == ErrEmptyName, err == ErrEmptyCategory, err == ErrEmptyPrice, err == ErrEmptyStock:
			payload = ResponseError{
				Success: false,
				Message: "ERR BAD REQUEST",
				Error:   err.Error(),
			}
			httpCode = http.StatusBadRequest
		case strings.Contains(err.Error(), "record not found"):
			payload = ResponseError{
				Success: false,
				Message: "ERR NOT FOUND",
				Error:   "data tidak ditemukan",
			}
		default:
			payload = ResponseError{
				Success: false,
				Message: "ERR INTERNAL",
				Error:   err.Error(),
			}
			httpCode = http.StatusInternalServerError
		}
		return c.Status(httpCode).JSON(payload)
	}
	return c.Status(fiber.StatusOK).JSON(ResponseSuccess{
		Success: true,
		Message: "UPDATE SUCCESS",
	})
}

func (h Handler) DeleteProduct(c *fiber.Ctx) error {
	idStr := c.Params("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(ResponseError{
			Success: false,
			Message: "ERR BAD REQUEST",
			Error:   "invalid id",
		})

	}

	err = h.svc.DeleteProduct(c.UserContext(), id)
	if err != nil {
		var payload ResponseError
		httpCode := 400
		switch {
		case strings.Contains(err.Error(), "record not found"):
			payload = ResponseError{
				Success: false,
				Message: "ERR NOT FOUND",
				Error:   "data tidak ditemukan",
			}
		default:
			payload = ResponseError{
				Success: false,
				Message: "ERR INTERNAL",
				Error:   err.Error(),
			}
			httpCode = http.StatusInternalServerError
		}
		return c.Status(httpCode).JSON(payload)
	}
	return c.Status(fiber.StatusOK).JSON(ResponseSuccess{
		Success: true,
		Message: "DELETE DATA SUCCESS",
	})
}
