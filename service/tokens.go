package service

import (
	"log"
	"time"
	"tokens-overhead/models"
	"tokens-overhead/repository"
)

type TokenService interface {
	GenerateJWT(numRoles int) (*string, error)
	Validate(token string) (bool, error)
	SleepAndValidate(token string, timeToSleepMs int) (bool, error)
	Execute(numRoles int, address string) error
}

type tokenServiceImpl struct {
	token            repository.TokenRepository
	roles            repository.RolesRepository
	request          repository.RequestInterface
	db               repository.DatabaseInterface
	machineName      string
	tokenCryptMethod string
}

func NewTokenService(
	roles repository.RolesRepository,
	token repository.TokenRepository,
	req repository.RequestInterface,
	db repository.DatabaseInterface,
	machineName, tokenCryptMethod string,
) TokenService {
	return &tokenServiceImpl{
		roles:            roles,
		token:            token,
		request:          req,
		db:               db,
		machineName:      machineName,
		tokenCryptMethod: tokenCryptMethod,
	}
}

func (t *tokenServiceImpl) GenerateJWT(numRoles int) (*string, error) {
	roles, err := t.roles.GetRoles(0, numRoles)
	if err != nil {
		return nil, err
	}
	// start := time.Now()
	jwt, err := t.token.Generate(roles)
	// log.Printf("JWT generation took %s for %5d roles and size=%dbytes", time.Since(start), numRoles, len(jwt))
	if err != nil {
		return nil, err
	}
	return &jwt, nil
}

func (t *tokenServiceImpl) Validate(token string) (bool, error) {
	return t.token.Validate(token)
}

func (t *tokenServiceImpl) SleepAndValidate(token string, timeToSleepMs int) (bool, error) {
	time.Sleep(time.Duration(timeToSleepMs) * time.Millisecond)

	return t.token.Validate(token)
}

func (t *tokenServiceImpl) Execute(numRoles int, address string) error {
	timeStartGeneratingToken := time.Now()
	token, err := t.GenerateJWT(numRoles)

	if err != nil {
		log.Printf("error generating token with %d roles: %s", numRoles, err)
		return err
	}
	timeGeneratingToken := time.Since(timeStartGeneratingToken)

	timeStartRequest := time.Now()
	err = t.request.Request(*token, address)
	if err != nil {
		log.Printf("error during request: %s", err)
		return err
	}
	timeRequesting := time.Since(timeStartRequest)

	return t.db.Save(
		models.NewRequest(
			timeRequesting.Microseconds(),
			t.machineName, address,
			len(*token), numRoles,
			t.tokenCryptMethod,
			timeGeneratingToken.Microseconds()))
}
