package user

// Service define la lógica de negocio relacionada a los usuarios.
type Service interface {
	Register(user *User) error
	Login(email, password string) (*User, error)
	VerifyKYC(userID int64, documentURLs []string) error
}
