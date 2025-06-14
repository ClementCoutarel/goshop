package database

type AuthService struct {
	Repo AuthRepo
}

func NewAuthService(repo AuthRepo) *AuthService {
	return &AuthService{
		Repo: *repo,
	}
}
