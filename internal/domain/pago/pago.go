package pago

import "time"

// Pago representa una transacción de pago en el sistema.
type Pago struct {
	ID            int64     `json:"id"`
	TransaccionID int64     `json:"transaccion_id"`
	Monto         float64   `json:"monto"`
	MetodoPago    string    `json:"metodo_pago"` // Ej. "Stripe", "PayPal", "criptomonedas"
	EstadoPago    string    `json:"estado_pago"` // Ej. "pendiente", "completado", "fallido"
	FechaPago     time.Time `json:"fecha_pago"`
}
