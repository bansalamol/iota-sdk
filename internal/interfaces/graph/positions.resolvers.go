package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.55

import (
	"context"
	"fmt"

	"github.com/iota-agency/iota-erp/internal/domain/entities/position"
	model "github.com/iota-agency/iota-erp/internal/interfaces/graph/gqlmodels"
	"github.com/iota-agency/iota-erp/sdk/mapper"
)

// CreatePosition is the resolver for the createPosition field.
func (r *mutationResolver) CreatePosition(ctx context.Context, input model.CreatePosition) (*model.Position, error) {
	entity := &position.Position{}
	if err := mapper.LenientMapping(&input, entity); err != nil {
		return nil, err
	}
	if err := r.app.PositionService.Create(ctx, entity); err != nil {
		return nil, err
	}
	return entity.ToGraph(), nil
}

// UpdatePosition is the resolver for the updatePosition field.
func (r *mutationResolver) UpdatePosition(ctx context.Context, id int64, input model.UpdatePosition) (*model.Position, error) {
	entity, err := r.app.PositionService.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if err := mapper.LenientMapping(&input, entity); err != nil {
		return nil, err
	}
	if err := r.app.PositionService.Update(ctx, entity); err != nil {
		return nil, err
	}
	return entity.ToGraph(), nil
}

// DeletePosition is the resolver for the deletePosition field.
func (r *mutationResolver) DeletePosition(ctx context.Context, id int64) (bool, error) {
	if err := r.app.PositionService.Delete(ctx, id); err != nil {
		return false, err
	}
	return true, nil
}

// Data is the resolver for the data field.
func (r *paginatedPositionsResolver) Data(ctx context.Context, obj *model.PaginatedPositions) ([]*model.Position, error) {
	entities, err := r.app.PositionService.GetPaginated(ctx, len(obj.Data), 0, nil)
	if err != nil {
		return nil, err
	}
	result := make([]*model.Position, len(entities))
	for i, entity := range entities {
		result[i] = entity.ToGraph()
	}
	return result, nil
}

// Total is the resolver for the total field.
func (r *paginatedPositionsResolver) Total(ctx context.Context, obj *model.PaginatedPositions) (int64, error) {
	return r.app.PositionService.Count(ctx)
}

// Position is the resolver for the position field.
func (r *queryResolver) Position(ctx context.Context, id int64) (*model.Position, error) {
	entity, err := r.app.PositionService.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return entity.ToGraph(), nil
}

// Positions is the resolver for the positions field.
func (r *queryResolver) Positions(ctx context.Context, offset int, limit int, sortBy []string) (*model.PaginatedPositions, error) {
	return &model.PaginatedPositions{}, nil
}

// PositionCreated is the resolver for the positionCreated field.
func (r *subscriptionResolver) PositionCreated(ctx context.Context) (<-chan *model.Position, error) {
	panic(fmt.Errorf("not implemented: PositionCreated - positionCreated"))
}

// PositionUpdated is the resolver for the positionUpdated field.
func (r *subscriptionResolver) PositionUpdated(ctx context.Context) (<-chan *model.Position, error) {
	panic(fmt.Errorf("not implemented: PositionUpdated - positionUpdated"))
}

// PositionDeleted is the resolver for the positionDeleted field.
func (r *subscriptionResolver) PositionDeleted(ctx context.Context) (<-chan int64, error) {
	panic(fmt.Errorf("not implemented: PositionDeleted - positionDeleted"))
}

// PaginatedPositions returns PaginatedPositionsResolver implementation.
func (r *Resolver) PaginatedPositions() PaginatedPositionsResolver {
	return &paginatedPositionsResolver{r}
}

type paginatedPositionsResolver struct{ *Resolver }
