package carga

import "time"

// Carga representa una carga publicada por un importador.
type Carga struct {
	ID            int64     `json:"id"`
	UserID        int64     `json:"user_id"` // Referencia al importador.
	Peso          float64   `json:"peso"`
	Volumen       float64   `json:"volumen"`
	TipoMercancia string    `json:"tipo_mercancia"`
	PaisDestino   string    `json:"pais_destino"`
	Estado        string    `json:"estado"` // Ej. "publicada", "en_transito", "entregada"
	CreatedAt     time.Time `json:"created_at"`
}
