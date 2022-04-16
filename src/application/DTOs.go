package application

import (
	"net/http"
)

type Product struct {
	Name string `json:"name"`
	Obs  string `json:"obs"`
}

func (a *Product) Bind(r *http.Request) error {
	return nil
}

type Result struct {
	StatusCod int    `json:"statusCod"`
	Message   string `json:"message"`
}
