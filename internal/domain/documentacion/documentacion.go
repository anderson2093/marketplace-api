package documentacion

import "time"

// Documentacion representa un documento asociado a un usuario o a una transacción,
// por ejemplo, documentos KYC o de embarque.
type Documentacion struct {
	ID            int64     `json:"id"`
	UsuarioID     int64     `json:"usuario_id"`               // Para documentos de verificación (KYC)
	TransaccionID int64     `json:"transaccion_id,omitempty"` // Opcional: para documentos asociados a una transacción
	Tipo          string    `json:"tipo"`                     // Ej. "KYC", "embarque"
	URLDocumento  string    `json:"url_documento"`            // Ruta o URL del archivo
	CreatedAt     time.Time `json:"created_at"`
}
