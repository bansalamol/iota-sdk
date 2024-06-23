package employee

import "context"

type Repository interface {
	Count(ctx context.Context) (int64, error)
	GetAll(ctx context.Context) ([]*Employee, error)
	GetPaginated(ctx context.Context, limit, offset int, sortBy []string) ([]*Employee, error)
	GetByID(ctx context.Context, id int64) (*Employee, error)
	GetEmployeeMeta(ctx context.Context, id int64) (*Meta, error)
	Create(ctx context.Context, upload *Employee) error
	Update(ctx context.Context, upload *Employee) error
	Delete(ctx context.Context, id int64) error
}
