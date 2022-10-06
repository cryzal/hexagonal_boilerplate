package outlet

type CreateReq struct {
	Name  string `json:"name,omitempty"`
	Email string `json:"email"`
	Phone string `json:"phone"`
}

type UpdateReq struct {
	Name  string `json:"name,omitempty"`
	Email string `json:"email"`
	Phone string `json:"phone"`
}
