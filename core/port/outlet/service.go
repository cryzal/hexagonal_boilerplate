package outlet

type Outlet struct {
	ID    *string
	Name  string
	Email string
	Phone string
}

type InportGetReq struct {
	OutletID string
}
type InportGetResp struct {
	Outlet Outlet
}
type InportCreateReq struct {
	Outlet Outlet
}
type InportUpdateReq struct {
	Outlet Outlet
}

type Service interface {
	Get(req InportGetReq) (InportGetResp, error)
	Create(req *InportCreateReq) error
	Update(req *InportUpdateReq) error
	UpdateWithPublisher(req *InportUpdateReq) error
}

func (cr *InportCreateReq) SetID(id *string) {
	cr.Outlet.ID = id
}
