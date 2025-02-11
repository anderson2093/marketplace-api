package evaluacion

// Repository define las operaciones de persistencia para las evaluaciones.
type Repository interface {
	Create(evaluacion *Evaluacion) error
	GetByID(id int64) (*Evaluacion, error)
	Update(evaluacion *Evaluacion) error
	Delete(id int64) error
}
