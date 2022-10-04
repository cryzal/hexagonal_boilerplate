package outlet

type (
	InportGetReq struct {
		OutletID string
	}
	Service interface {
		Get(req InportGetReq) error
	}
)
