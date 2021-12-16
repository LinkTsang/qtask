package service

import "context"

type Service interface {
	Health(ctx context.Context) (bool, error)
}
