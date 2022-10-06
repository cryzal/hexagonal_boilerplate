package service

import (
	"context"
	"hexagonal_boilerplate/core/entities"
	port "hexagonal_boilerplate/core/port/outlet"
	portPublisher "hexagonal_boilerplate/core/port/publisher"
	portDBTransaction "hexagonal_boilerplate/core/port/transaction"
	"hexagonal_boilerplate/shared/models/event"
)

type Service struct {
	OutletRepo      port.Repository
	TransactionRepo portDBTransaction.Repository
	PublisherImpl   portPublisher.Publisher
}

func OutletServiceNew(OutletRepo port.Repository, TransactionRepo portDBTransaction.Repository, PublisherImpl portPublisher.Publisher) *Service {
	return &Service{OutletRepo, TransactionRepo, PublisherImpl}
}

func (s *Service) Get(req port.InportGetReq) (port.InportGetResp, error) {
	var resp port.InportGetResp
	outlet, err := s.OutletRepo.Get(req.OutletID)

	if err != nil {
		return resp, err
	}
	resp.Outlet = port.Outlet{
		ID:    outlet.Code,
		Name:  outlet.Name,
		Email: outlet.Email,
		Phone: outlet.Phone,
	}
	return resp, nil
}

func (s *Service) Create(req *port.InportCreateReq) error {

	outlet, err := entities.NewOutlet(req.Outlet.Name, req.Outlet.Email, req.Outlet.Phone)
	if err != nil {
		return err
	}

	c := context.Background()
	err = s.TransactionRepo.WithinTransaction(c, func(ctx context.Context) error {
		err := s.OutletRepo.Create(ctx, outlet)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return err
	}

	req.SetID(outlet.Code)
	return nil
}

func (s *Service) Update(req *port.InportUpdateReq) error {
	outlet, err := s.OutletRepo.Get(*req.Outlet.ID)
	if err != nil {
		return err
	}
	reqEditOutlet := entities.ReqEditOutlet{
		Name:  req.Outlet.Name,
		Email: req.Outlet.Email,
		Phone: req.Outlet.Phone,
	}
	err = outlet.EditOutlet(reqEditOutlet)
	if err != nil {
		return err
	}

	c := context.Background()

	err = s.TransactionRepo.WithinTransaction(c, func(ctx context.Context) error {
		err := s.OutletRepo.Update(ctx, outlet)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) UpdateWithPublisher(req *port.InportUpdateReq) error {
	err := s.PublisherImpl.PublishMessage(event.UpdateOutletEvent, req.Outlet)
	if err != nil {
		return err
	}
	return nil
}
