package handler

import (
	"net/http"
	"strconv"

	// "pt-xyz-multifinance/internal/domain"

	"pt-xyz-multifinance/internal/service"

	"github.com/labstack/echo/v4"
)

type LimitHandler struct {
	service *service.LimitService
}

func NewLimitHandler(service *service.LimitService) *LimitHandler {
	return &LimitHandler{service: service}
}

func (h *LimitHandler) GetConsumerLimit(c echo.Context) error {

	id, err := strconv.ParseUint(c.Param("consumerId"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid Consumer ID"})
	}

	limits, err := h.service.GetConsumerLimit(uint(id))
	// fmt.Print("Received consumerId: ", limits)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, limits)
}

func (h *LimitHandler) CreateConsumerLimits(c echo.Context) error {
	// log.Printf("Received consumerId: %s", c.Param("consumerId"))

	consumerID, err := strconv.ParseUint(c.Param("consumerId"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid consumer ID",
		})
	}

	// Proses penetapan limit
	limits, err := h.service.CreateLimitsForConsumer(uint(consumerID))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, limits)
}
