package outlet

import (
	"fmt"
	port "hexagonal_boilerplate/core/port/outlet"
	"hexagonal_boilerplate/interface/rest/common/response"
	"hexagonal_boilerplate/shared/messaging"
	logKey "hexagonal_boilerplate/shared/utils/logger/libs/const/log"
	"hexagonal_boilerplate/shared/utils/logger/log"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

type Handler struct {
	Service port.Service
}

func New(Service port.Service) *Handler {
	return &Handler{Service}
}

func (h *Handler) Create(c echo.Context) error {
	startTime := time.Now()
	request := CreateReq{}
	resp := response.Response{}
	data := CreateResp{}
	err := c.Bind(&request)
	if err != nil {
		return c.JSON(http.StatusBadRequest, nil)
	}

	log := log.With().
		Interface(logKey.REQUEST_PAYLOAD, request).
		Str(logKey.HANDLER_NAME, logKey.OUTLET_HANDLER).
		Str(logKey.HANDLER_METHOD, "Create").
		Time(logKey.START_TIME, startTime).
		Logger()
	log.Info().Msg("")

	req := port.InportCreateReq{
		Outlet: port.Outlet{
			Name:  request.Name,
			Email: request.Email,
			Phone: request.Phone,
		},
	}
	err = h.Service.Create(&req)
	if err != nil {
		log.Error().
			Dur(logKey.LATENCY, time.Since(startTime)).
			Msg(err.Error())
		resp.RespCode = http.StatusBadRequest
		resp.Message = err.Error()
		return c.JSON(http.StatusBadRequest, resp.GetResponsePayloadIntl())
	}
	data.ID = *req.Outlet.ID

	resp.RespCode = http.StatusCreated
	resp.Message = "Success Storing Data"
	resp.Product = data

	log.Info().
		Interface(logKey.RESPONSE_PAYLOAD, resp.GetResponsePayloadIntl()).
		Dur(logKey.LATENCY, time.Since(startTime)).
		Msg(logKey.OK)
	return c.JSON(200, resp.GetResponsePayloadIntl())
}

func (h *Handler) Get(c echo.Context) error {
	startTime := time.Now()
	resp := response.Response{}
	OutletId := c.Param("outlet_id")

	log := log.With().
		Interface(logKey.OUTLET_ID, OutletId).
		Str(logKey.HANDLER_NAME, logKey.OUTLET_HANDLER).
		Str(logKey.HANDLER_METHOD, "Get").
		Time(logKey.START_TIME, startTime).
		Logger()
	log.Info().Msg("")

	req := port.InportGetReq{
		OutletID: OutletId,
	}
	result, err := h.Service.Get(req)

	if err != nil {
		log.Error().
			Dur(logKey.LATENCY, time.Since(startTime)).
			Msg(err.Error())
		resp.RespCode = http.StatusBadRequest
		resp.Message = err.Error()
		return c.JSON(http.StatusBadRequest, resp.GetResponsePayloadIntl())
	}

	resp.RespCode = http.StatusOK
	resp.Message = "Success Retrieving Data"
	resp.Outlet = result
	log.Info().
		Interface(logKey.RESPONSE_PAYLOAD, resp.GetResponsePayloadIntl()).
		Dur(logKey.LATENCY, time.Since(startTime)).
		Msg(logKey.OK)
	return c.JSON(http.StatusOK, resp.GetResponsePayloadIntl())
}
func (h *Handler) Edit(c echo.Context) error {
	request := UpdateReq{}

	resp := response.Response{}
	OutletId := c.Param("outlet_id")
	err := c.Bind(&request)
	if err != nil {
		return c.JSON(http.StatusBadRequest, nil)
	}
	req := port.InportUpdateReq{
		Outlet: port.Outlet{
			ID:    &OutletId,
			Name:  request.Name,
			Email: request.Email,
			Phone: request.Phone,
		},
	}
	err = h.Service.Update(&req)
	if err != nil {
		resp.RespCode = http.StatusBadRequest
		resp.Message = err.Error()
		return c.JSON(http.StatusBadRequest, resp.GetResponsePayloadIntl())
	}
	// passing data to response contract
	resp.RespCode = http.StatusNoContent
	resp.Message = "Success Update Data"
	return c.JSON(http.StatusNoContent, resp.GetResponsePayloadIntl())
}
func (h *Handler) EditPublisher(c echo.Context) error {
	request := UpdateReq{}

	resp := response.Response{}
	OutletId := c.Param("outlet_id")
	err := c.Bind(&request)
	if err != nil {
		return c.JSON(http.StatusBadRequest, nil)
	}
	req := port.InportUpdateReq{
		Outlet: port.Outlet{
			ID:    &OutletId,
			Name:  request.Name,
			Email: request.Email,
			Phone: request.Phone,
		},
	}
	err = h.Service.UpdateWithPublisher(&req)
	if err != nil {
		resp.RespCode = http.StatusBadRequest
		resp.Message = err.Error()
		return c.JSON(http.StatusBadRequest, resp.GetResponsePayloadIntl())
	}
	// passing data to response contract
	resp.RespCode = http.StatusNoContent
	resp.Message = "Success Update Data"
	return c.JSON(http.StatusNoContent, resp.GetResponsePayloadIntl())
}

func (h *Handler) Update2() messaging.HandleFunc {
	return func(payload messaging.Payload, err error) {
		fmt.Println(payload.Data)
	}
}
