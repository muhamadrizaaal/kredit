package handler

import (
	"net/http"
	"strconv"

	"pt-xyz-multifinance/internal/domain"
	"pt-xyz-multifinance/internal/service"

	"github.com/labstack/echo/v4"
)

type ConsumerHandler struct {
	service *service.ConsumerService
}

func NewConsumerHandler(service *service.ConsumerService) *ConsumerHandler {
	return &ConsumerHandler{service: service}
}

func (h *ConsumerHandler) CreateConsumer(c echo.Context) error {
	var consumer domain.Consumer
	if err := c.Bind(&consumer); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	if err := h.service.CreateConsumer(&consumer); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusCreated, consumer)
}

func (h *ConsumerHandler) GetConsumer(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid ID"})
	}

	consumer, err := h.service.GetConsumer(uint(id))
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Consumer not found"})
	}

	return c.JSON(http.StatusOK, consumer)
}
