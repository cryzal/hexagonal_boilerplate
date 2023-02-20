package webhook

import (
	"github.com/labstack/echo/v4"
	"hexagonal_boilerplate/core/port/webhook"
	"net/http"
)

type Handler struct {
	port webhook.Service
}

func Newhandler(port webhook.Service) *Handler {
	return &Handler{
		port: port,
	}
}
func (h *Handler) Shopee(c echo.Context) error {
	req := ShopeeReq{}

	err := c.Bind(&req)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err)
	}

	err = h.port.Order("SHOPEE", req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	return nil
}
