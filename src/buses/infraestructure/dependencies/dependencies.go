package dependencies

import (
	"consumer/core"
	"consumer/src/buses/application"
	"consumer/src/buses/infraestructure"
	"consumer/src/buses/infraestructure/controllers"
)

var (
	mySQL infraestructure.MySQL
)

func InitBus() {
	db,err := core.InitMySQL()
	if err != nil {
		return
	}
	mySQL = *infraestructure.NewMySQL(db)
}


func UpdateBusController() *controllers.UpdateBusByIDController {
	ucUpdateBus := application.NewUpdateBusByIDUseCase(&mySQL)

	return controllers.NewUpdateBusByIDController(ucUpdateBus)
}
