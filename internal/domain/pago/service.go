package pago

// Service define la lógica de negocio para el procesamiento de pagos.
type Service interface {
	ProcessPayment(transaccionID int64, monto float64, metodo string) (*Pago, error)
	Refund(pagoID int64) error
}
