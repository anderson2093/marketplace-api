package transaccion

// Repository define las operaciones de persistencia para las transacciones.
type Repository interface {
	Create(transaccion *Transaccion) error
	GetByID(id int64) (*Transaccion, error)
	Update(transaccion *Transaccion) error
	Delete(id int64) error
}
