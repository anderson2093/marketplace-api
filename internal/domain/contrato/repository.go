package contrato

// Repository define las operaciones de persistencia para los contratos.
type Repository interface {
	Create(contrato *Contrato) error
	GetByID(id int64) (*Contrato, error)
	Update(contrato *Contrato) error
	Delete(id int64) error
}
