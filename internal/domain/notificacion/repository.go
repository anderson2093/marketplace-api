package notificacion

// Repository define las operaciones de persistencia para las notificaciones.
type Repository interface {
	Create(notificacion *Notificacion) error
	GetByID(id int64) (*Notificacion, error)
	Update(notificacion *Notificacion) error
	Delete(id int64) error
}
