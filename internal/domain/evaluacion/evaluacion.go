package evaluacion

import "time"

// Evaluacion representa la valoración de una transacción por parte de un usuario.
type Evaluacion struct {
	ID            int64     `json:"id"`
	TransaccionID int64     `json:"transaccion_id"`
	UsuarioID     int64     `json:"usuario_id"`
	Puntuacion    int       `json:"puntuacion"` // Por ejemplo, de 1 a 5.
	Comentario    string    `json:"comentario"`
	CreatedAt     time.Time `json:"created_at"`
}
