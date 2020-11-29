package service

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/nicobianchetti/Go-DemoConcurrencia/model"
)

//ICarDetailsService implements methods for Car Details
type ICarDetailsService interface {
	GetDetails() *model.CarDetails
}

var (
	carService       ICarService   = NewCarService()
	ownerService     IOwnerService = NewOwnerService()
	carDataChannel                 = make(chan *http.Response)
	ownerDataChannel               = make(chan *http.Response)
)

type service struct{}

//NewCarDetailsService instance of
func NewCarDetailsService() ICarDetailsService {
	return &service{}
}

func (s *service) GetDetails() *model.CarDetails {
	//goroutine get data from https://myfakeapi.com/api/cars/1 --> endpoint 1
	go carService.FetchDataCar()

	//goroutine get data from https://myfakeapi.com/api/users/1
	go ownerService.FetchDataOwner()

	car, _ := getCarData()
	owner, _ := getOwnerData()

	//create carChannel to get the data from endpoint2
	//create ownerChannel to ghe the data from endpoint2

	return &model.CarDetails{
		ID:        car.ID,
		Brand:     car.Brand,
		Model:     car.Model,
		Year:      car.Year,
		FirstName: owner.FirstName,
		LastName:  owner.LastName,
		Email:     owner.Email,
	}
}

func getCarData() (model.Car, error) {
	r1 := <-carDataChannel
	var car model.Car

	err := json.NewDecoder(r1.Body).Decode(&car)

	if err != nil {
		fmt.Print(err.Error())
		return car, err
	}

	return car, nil
}

func getOwnerData() (model.Owner, error) {
	r1 := <-ownerDataChannel
	var owner model.Owner

	err := json.NewDecoder(r1.Body).Decode(&owner)

	if err != nil {
		fmt.Print(err.Error())
		return owner, err
	}

	return owner, nil
}
