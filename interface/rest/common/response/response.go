package response

import (
	"strconv"
)

type Response struct {
	Message          string
	RespCode         int
	Outlet           interface{}
	Outlets          interface{}
	Meta             interface{}
	Breadcrumb       interface{}
	Breadcrumbs      interface{}
	Product          interface{}
	Products         interface{}
	ErrorsValidation interface{}
}

type ResponseIntl struct {
	ResponseCode     string      `json:"response_code"`
	Message          string      `json:"message"`
	Outlet           interface{} `json:"outlet,omitempty"`
	ErrorsValidation interface{} `json:"errors_validation,omitempty"`
	Outlets          interface{} `json:"outlets,omitempty"`
	Meta             interface{} `json:"meta,omitempty"`
	Breadcrumb       interface{} `json:"breadcrumb,omitempty"`
	Breadcrumbs      interface{} `json:"breadcrumbs,omitempty"`
	Product          interface{} `json:"product,omitempty"`
	Products         interface{} `json:"products,omitempty"`
}

func (res *Response) GetResponsePayloadIntl() ResponseIntl {
	resp := ResponseIntl{}
	resp.ResponseCode = strconv.Itoa(res.RespCode)
	resp.Message = res.Message
	resp.Outlet = res.Outlet
	resp.ErrorsValidation = res.ErrorsValidation
	resp.Outlets = res.Outlets
	resp.Meta = res.Meta
	resp.Breadcrumb = res.Breadcrumb
	resp.Breadcrumbs = res.Breadcrumbs
	resp.Product = res.Product
	resp.Products = res.Products
	return resp
}
