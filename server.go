package main

import (
	"github.com/nicobianchetti/Go-DemoConcurrencia/controller"
	"github.com/nicobianchetti/Go-DemoConcurrencia/router"
	"github.com/nicobianchetti/Go-DemoConcurrencia/service"
)

var (
	carDetailsService    service.ICarDetailsService      = service.NewCarDetailsService()
	carDetailsController controller.CarDetailsController = controller.NewCarDetailsController(carDetailsService)
	httpRouter           router.IRouter                  = router.NewChiRouter()
)

func main() {
	const port string = ":8080"
	httpRouter.GET("/carDetails", carDetailsController.GetCarDetails)
	httpRouter.SERVE(port)
}
