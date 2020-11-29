package service

import (
	"fmt"
	"net/http"
)

//IOwnerService implements methods for owner-service
type IOwnerService interface {
	FetchDataOwner()
}

const (
	ownerServiceURL = "https://myfakeapi.com/api/cars/1"
)

type fetchOwnerDataService struct {
}

//NewOwnerService instance a new OwnerService
func NewOwnerService() IOwnerService {
	return &fetchCarDataService{}
}

func (f *fetchCarDataService) FetchDataOwner() {
	client := http.Client{}

	fmt.Printf("Fectching the irl %s", ownerServiceURL)

	//Call the external API
	resp, _ := client.Get(ownerServiceURL)

	//Write response to the channel (TODO) (el channel necesita ser compartido por los dos services de car)
	ownerDataChannel <- resp
}
