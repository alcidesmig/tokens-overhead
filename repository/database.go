package repository

import "tokens-overhead/models"

type DatabaseInterface interface {
	Save(r models.Request) error
}
