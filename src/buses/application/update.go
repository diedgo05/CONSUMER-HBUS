package application

import "consumer/src/buses/domain"

type UpdateBusByIDUseCase struct {
	db domain.IBusesRepository
}

func NewUpdateBusByIDUseCase(db domain.IBusesRepository) *UpdateBusByIDUseCase {
	return &UpdateBusByIDUseCase{db: db}
}

func (uc *UpdateBusByIDUseCase) Run(idBus int, bus domain.Buses) error {
	return uc.db.UpdateByID(idBus, bus)
}