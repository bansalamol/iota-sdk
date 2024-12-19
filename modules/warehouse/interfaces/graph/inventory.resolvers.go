package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.57

import (
	"context"
	"fmt"

	model "github.com/iota-agency/iota-sdk/modules/warehouse/interfaces/graph/gqlmodels"
)

// CompleteInventoryCheck is the resolver for the completeInventoryCheck field.
func (r *mutationResolver) CompleteInventoryCheck(ctx context.Context, items []*model.InventoryItem) (*model.InventoryPosition, error) {
	panic(fmt.Errorf("not implemented: CompleteInventoryCheck - completeInventoryCheck"))
}

// InventoryPositions is the resolver for the inventoryPositions field.
func (r *queryResolver) InventoryPositions(ctx context.Context, offset int, limit int, sortBy []string) ([]*model.InventoryPosition, error) {
	panic(fmt.Errorf("not implemented: InventoryPositions - inventoryPositions"))
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
