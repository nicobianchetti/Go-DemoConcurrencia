package controller

import (
	"encoding/json"
	"net/http"

	"github.com/nicobianchetti/Go-DemoConcurrencia/service"
)

var (
	carDetailsService service.ICarDetailsService
)

type CarDetailsController interface {
	GetCarDetails(w http.ResponseWriter, r *http.Request)
}

type controller struct {
}

func NewCarDetailsController(service service.ICarDetailsService) CarDetailsController {
	carDetailsService = service
	return &controller{}
}

func (c *controller) GetCarDetails(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	result := carDetailsService.GetDetails()
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)
}
