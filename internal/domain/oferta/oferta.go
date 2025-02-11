package oferta

import "time"

// OfertaContenedor representa una oferta de espacio en un contenedor publicada por un exportador.
type OfertaContenedor struct {
	ID                int64     `json:"id"`
	UserID            int64     `json:"user_id"` // Referencia al exportador.
	CapacidadTotal    float64   `json:"capacidad_total"`
	VolumenDisponible float64   `json:"volumen_disponible"`
	PesoDisponible    float64   `json:"peso_disponible"`
	Costo             float64   `json:"costo"`
	FechaEmbarque     time.Time `json:"fecha_embarque"`
	Ruta              string    `json:"ruta"`   // Ej. "origen-destino"
	Estado            string    `json:"estado"` // Ej. "disponible", "reservado"
	CreatedAt         time.Time `json:"created_at"`
}
