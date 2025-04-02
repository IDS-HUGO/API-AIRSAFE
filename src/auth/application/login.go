package application

import (
	"apiusersafe/src/auth/domain"
)

type LoginService struct {
	UserRepo   domain.UserRepository
	JWTService domain.JWTService
}

type LoginResponse struct {
	User  *domain.User `json:"user"`
	Token string       `json:"token"`
}

func NewLoginService(userRepo domain.UserRepository, jwtService domain.JWTService) *LoginService {
	return &LoginService{
		UserRepo:   userRepo,
		JWTService: jwtService,
	}
}

func (s *LoginService) Execute(usuario, contrasena string) (*LoginResponse, error) {
	user, err := s.UserRepo.ValidateUser(usuario, contrasena)
	if err != nil {
		return nil, err
	}

	token, err := s.JWTService.GenerateToken(user.Usuario)
	if err != nil {
		return nil, err
	}

	return &LoginResponse{
		User:  user,
		Token: token,
	}, nil
}
