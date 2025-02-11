package oferta

// Repository define las operaciones de persistencia para las ofertas de contenedor.
type Repository interface {
	Create(oferta *OfertaContenedor) error
	GetByID(id int64) (*OfertaContenedor, error)
	Update(oferta *OfertaContenedor) error
	Delete(id int64) error
}
