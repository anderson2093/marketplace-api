package notificacion

// Service define la lógica de negocio para el envío y manejo de notificaciones.
type Service interface {
	SendNotification(usuarioID int64, mensaje, tipo string) error
	MarkAsRead(notificacionID int64) error
}
