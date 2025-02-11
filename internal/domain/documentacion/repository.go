package documentacion

// Repository define las operaciones de persistencia para la documentación.
type Repository interface {
	Create(documentacion *Documentacion) error
	GetByID(id int64) (*Documentacion, error)
	Update(documentacion *Documentacion) error
	Delete(id int64) error
}
