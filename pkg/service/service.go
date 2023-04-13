package service

import (
	"errors"
	"strings"
)

var ErrEmpty = errors.New("empty string")

type StringService interface {
	Uppercase(string) (string, error)
	Count(string) int
}

type stringService struct{}

func New() StringService {
	return &stringService{}
}

func (ss *stringService) Uppercase(s string) (string, error) {
	if s == "" {
		return "", ErrEmpty
	}
	return strings.ToUpper(s), nil
}

func (ss *stringService) Count(s string) int {
	return len(s)
}
