package repository

//go:generate mockgen -destination=mock/token_repository.go -package=mock . TokenRepository
type TokenRepository interface {
	Generate(roles []string) (string, error)
	Validate(token string) (bool, error)
}
