package domain

type JWTService interface {
	GenerateToken(usuario string) (string, error)
	ValidateToken(token string) (string, error)
}
