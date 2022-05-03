package froles

import (
	"errors"
	"log"
	"os"
	"strings"
	"tokens-overhead/repository"
)

type fileRolesImpl struct {
	roles []string
}

func NewFileRolesImpl(path string) (repository.RolesRepository, error) {
	f := fileRolesImpl{}
	if path != "" {
		err := f.LoadRoles(path)
		if err != nil {
			return nil, err
		}
	}
	return &f, nil
}

func (f *fileRolesImpl) LoadRoles(path string) error {
	content, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	f.roles = strings.Split(string(content), "\n")
	log.Printf("loaded %d roles", len(f.roles))
	return nil
}

func (f *fileRolesImpl) GetRoles(offset, limit int) ([]string, error) {
	if offset < 0 || offset > limit || offset > len(f.roles) {
		return nil, errors.New("invalid offset")
	}
	if limit < 0 || limit > len(f.roles) {
		return nil, errors.New("invalid limit")
	}
	return f.roles[offset:limit], nil
}
