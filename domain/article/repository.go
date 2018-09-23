package article

import "context"

type Repository interface {
	Create(ctx context.Context, entity *Entity) error
}
