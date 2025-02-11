package user

import "time"

// Role define el rol del usuario en el sistema.
type Role string

const (
	RoleImportador Role = "importador"
	RoleExportador Role = "exportador"
	RoleAgente     Role = "agente_logistico"
)

// User representa al usuario en el sistema.
type User struct {
	ID            int64     `json:"id"`
	Nombre        string    `json:"nombre"`
	Email         string    `json:"email"`
	Password      string    `json:"password"` // Se recomienda almacenar el hash.
	Role          Role      `json:"role"`
	KYCVerificado bool      `json:"kyc_verificado"`
	CreatedAt     time.Time `json:"created_at"`
}
