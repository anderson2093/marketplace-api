package evaluacion

// Service define la lógica de negocio para la gestión de evaluaciones.
type Service interface {
	SubmitEvaluation(evaluacion *Evaluacion) error
	GetEvaluationsByTransaccion(transaccionID int64) ([]*Evaluacion, error)
}
