package outlet

import (
	"encoding/json"
	port "hexagonal_boilerplate/core/port/outlet"
	"hexagonal_boilerplate/shared/messaging"
	"log"
)

type Handler struct {
	Service port.Service
}

func New(Service port.Service) *Handler {
	return &Handler{Service}
}

func (h *Handler) Update() messaging.HandleFunc {
	return func(payload messaging.Payload, err error) {
		dbByte, err := json.Marshal(payload.Data)
		if err != nil {
			log.Println(err.Error())
			return
		}
		var dataPayload Outlet
		err = json.Unmarshal(dbByte, &dataPayload)
		if err != nil {
			log.Println(err.Error())
			return
		}
		req := port.InportUpdateReq{
			Outlet: port.Outlet{
				ID:    dataPayload.ID,
				Name:  dataPayload.Name,
				Email: dataPayload.Email,
				Phone: dataPayload.Phone,
			},
		}
		log.Println(req)
		err = h.Service.Update(&req)
		if err != nil {
			log.Println(err.Error())
			return
		}
		return
	}
}
