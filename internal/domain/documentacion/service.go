package documentacion

// Service define la lógica de negocio para la gestión de documentos.
type Service interface {
	UploadDocument(documentacion *Documentacion) error
	ValidateDocument(documentacionID int64) error
}
