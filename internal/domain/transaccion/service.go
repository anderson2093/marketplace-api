package transaccion

// Service define la lógica de negocio para la consolidación de cargas.
type Service interface {
	Consolidate(cargaID, ofertaID int64, espacioAsignado, costoCompartido float64) (*Transaccion, error)
	UpdateStatus(id int64, newStatus string) error
}
