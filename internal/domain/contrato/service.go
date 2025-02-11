package contrato

// Service define la lógica de negocio para la generación y validación de contratos.
type Service interface {
	Generate(transaccionID int64, detalles string) (*Contrato, error)
	Validate(contratoID int64) error
}
