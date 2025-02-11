package notificacion

import "time"

// Notificacion representa una alerta o actualización enviada a un usuario.
type Notificacion struct {
	ID         int64     `json:"id"`
	UsuarioID  int64     `json:"usuario_id"`
	Mensaje    string    `json:"mensaje"`
	Tipo       string    `json:"tipo"` // Ej. "info", "alerta"
	Leido      bool      `json:"leido"`
	FechaEnvio time.Time `json:"fecha_envio"`
}
