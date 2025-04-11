package infraestructure

import (
	"consumer/src/buses/domain"
	"database/sql"
	"fmt"
)

type MySQL struct {
	db *sql.DB
}

func NewMySQL(db *sql.DB) *MySQL {
	return &MySQL{db: db}
}

func (mysql *MySQL) UpdateByID(idBus int, bus domain.Buses) error {
	query := "UPDATE buses SET  disponible =? WHERE idBus = ?"
	_, err := mysql.db.Exec(query,  bus.Disponible, idBus)

	if err != nil {
		return err
	}

	fmt.Println("Datos del bus actualizados correctamente")
	return nil
}

