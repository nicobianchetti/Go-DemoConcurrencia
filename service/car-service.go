package service

import (
	"fmt"
	"net/http"
)

//ICarService implements methods for CarService
type ICarService interface {
	FetchDataCar()
}

const (
	carServiceURL = "https://myfakeapi.com/api/cars/1"
)

type fetchCarDataService struct {
}

//NewCarService instanced a new object of interface
func NewCarService() ICarService {
	return &fetchCarDataService{}
}

func (f *fetchCarDataService) FetchDataCar() {
	client := http.Client{}

	fmt.Printf("Fectching the irl %s", carServiceURL)

	//Call the external API
	resp, _ := client.Get(carServiceURL)

	//Write response to the channel  (el channel necesita ser compartido por los dos services de car,este (escribe) y car-details service (lee) )
	carDataChannel <- resp
}
