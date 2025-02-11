package contrato

import "time"

// Contrato formaliza el acuerdo entre las partes en una transacción.
type Contrato struct {
	ID            int64     `json:"id"`
	TransaccionID int64     `json:"transaccion_id"`
	Detalles      string    `json:"detalles"` // Condiciones, cláusulas, etc.
	FechaEmision  time.Time `json:"fecha_emision"`
	Estado        string    `json:"estado"` // Ej. "activo", "expirado"
}
