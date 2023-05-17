package repository

import (
	"context"
)

type StudentRepo interface {
	Add(ctx context.Context, student *Student) (uint64, error)
	GetById(ctx context.Context, id uint64) (*Student, error)
	Update(ctx context.Context, student *Student) (bool, error)
	Delete(ctx context.Context, id uint64) (bool, error)
}
