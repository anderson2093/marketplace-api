package pago

// Repository define las operaciones de persistencia para los pagos.
type Repository interface {
	Create(pago *Pago) error
	GetByID(id int64) (*Pago, error)
	Update(pago *Pago) error
	Delete(id int64) error
}
