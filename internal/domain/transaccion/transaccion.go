package transaccion

import "time"

// Transaccion representa la consolidación de una carga con una oferta de contenedor.
type Transaccion struct {
	ID              int64     `json:"id"`
	CargaID         int64     `json:"carga_id"`
	OfertaID        int64     `json:"oferta_id"`
	ContenedorID    int64     `json:"contenedor_id,omitempty"` // Opcional, si se asigna un contenedor físico.
	EspacioAsignado float64   `json:"espacio_asignado"`        // Volumen o peso asignado.
	CostoCompartido float64   `json:"costo_compartido"`
	Estado          string    `json:"estado"` // Ej. "pendiente", "confirmada", "completada", "cancelada"
	CreatedAt       time.Time `json:"created_at"`
}
