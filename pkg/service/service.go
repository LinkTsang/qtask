package service

type Service interface {
	Health() bool
}
