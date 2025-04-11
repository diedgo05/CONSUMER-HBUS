package domain

type IBusesRepository interface {
	UpdateByID(idBus int, bus Buses) error
}