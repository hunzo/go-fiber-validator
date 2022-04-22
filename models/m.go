package models

type RequestModels struct {
	UserName string `json:"username" validate:"required,min=6,max=10"`
	Password string `json:"password" validate:"required"`
	Roles    string `json:"roles" validate:"required"`
}

type ErrorResponse struct {
	FailedField string
	Tag         string
	Value       string
}
