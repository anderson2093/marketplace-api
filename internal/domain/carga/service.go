package carga

// Service define la lógica de negocio relacionada a las cargas.
type Service interface {
	Publish(carga *Carga) error
	UpdateStatus(id int64, newStatus string) error
}
