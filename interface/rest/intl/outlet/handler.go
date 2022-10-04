package outlet

import (
	"fmt"
	"hexagonal_boilerplate/shared/messaging"

	"github.com/labstack/echo/v4"

	port "hexagonal_boilerplate/core/port/outlet"
)

type Handler struct {
	App port.Service
}

func New(App port.Service) *Handler {
	return &Handler{App}
}

func (h *Handler) Get(c echo.Context) error {
	OutletId := c.Param("outlet_id")
	fmt.Println(OutletId)
	req := port.InportGetReq{
		OutletID: OutletId,
	}
	err := h.App.Get(req)
	if err != nil {
		return err
	}
	return c.JSON(200, "ok")
}

func (h *Handler) Update() messaging.HandleFunc {
	return func(payload messaging.Payload, err error) {
		fmt.Println(payload.Data)
	}
}

func (h *Handler) Update2() messaging.HandleFunc {
	return func(payload messaging.Payload, err error) {
		fmt.Println(payload.Data)
	}
}
