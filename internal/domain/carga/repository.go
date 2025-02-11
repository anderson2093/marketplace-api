package carga

// Repository define las operaciones de persistencia para las cargas.
type Repository interface {
	Create(carga *Carga) error
	GetByID(id int64) (*Carga, error)
	Update(carga *Carga) error
	Delete(id int64) error
}
