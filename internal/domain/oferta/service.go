package oferta

// Service define la lógica de negocio relacionada a las ofertas de contenedor.
type Service interface {
	Publish(oferta *OfertaContenedor) error
	UpdateAvailability(id int64, availableVolume, availableWeight float64) error
}
