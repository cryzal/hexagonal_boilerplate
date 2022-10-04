package service

import (
	"fmt"
	Port "hexagonal_boilerplate/core/port/outlet"
)

type Service struct {
	OutletRepo Port.Repository
}

func OutletServiceNew(OutletRepo Port.Repository) *Service {
	return &Service{OutletRepo}
}

func (o *Service) Get(req Port.InportGetReq) error {
	fmt.Println(req)
	return nil
}
