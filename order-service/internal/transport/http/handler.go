package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"order-service/internal/domain"
)

type Handler struct {
	uc interface {
		CreateOrder(amount int64, email string) domain.CreateOrderResponse
	}
}

func NewHandler(u interface {
	CreateOrder(amount int64, email string) domain.CreateOrderResponse
}) *Handler {
	return &Handler{uc: u}
}

func (h *Handler) CreateOrder(c *gin.Context) {
	var req struct {
	Amount int64  `json:"amount"`
	Email  string `json:"email"`
}

	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	res := h.uc.CreateOrder(req.Amount, req.Email)

	c.JSON(http.StatusOK, res)
}