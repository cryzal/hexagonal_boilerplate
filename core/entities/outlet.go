package entities

import (
	"errors"
)

type Outlet struct {
	ID    *int64
	Code  *string
	Name  string
	Email string
	Phone string
}

type ReqEditOutlet struct {
	Name  string
	Email string
	Phone string
}

func NewOutlet(name, email, phone string) (*Outlet, error) {
	var obj Outlet
	obj.Name = name
	obj.Phone = phone
	obj.Email = email

	err := obj.Validate()

	if err != nil {
		return nil, err
	}

	return &obj, nil
}

func (o *Outlet) SetID(id *int64, code *string) {
	o.ID = id
	o.Code = code
}

func (o *Outlet) EditOutlet(req ReqEditOutlet) error {
	o.Name = req.Name
	o.Email = req.Email
	o.Phone = req.Phone
	err := o.Validate()

	if err != nil {
		return err
	}
	return nil
}

func (o *Outlet) Validate() error {
	if o.Phone == "" {
		return errors.New("phone is required")
	}
	if o.Email == "" {
		return errors.New("email is required")
	}
	if o.Name == "" {
		return errors.New("name is required")
	}
	return nil
}
