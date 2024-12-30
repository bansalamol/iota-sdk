package services

import (
	"context"

	"github.com/iota-uz/iota-sdk/modules/core/domain/entities/tab"
	"github.com/iota-uz/iota-sdk/pkg/composables"
)

type TabService struct {
	repo tab.Repository
}

func NewTabService(repo tab.Repository) *TabService {
	return &TabService{repo}
}

func (s *TabService) GetByID(ctx context.Context, id uint) (*tab.Tab, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *TabService) GetAll(ctx context.Context, params *tab.FindParams) ([]*tab.Tab, error) {
	return s.repo.GetAll(ctx, params)
}

func (s *TabService) GetUserTabs(ctx context.Context, userID uint) ([]*tab.Tab, error) {
	return s.repo.GetUserTabs(ctx, userID)
}

func (s *TabService) Create(ctx context.Context, data *tab.CreateDTO) (*tab.Tab, error) {
	tx, err := composables.UsePoolTx(ctx)
	if err != nil {
		return nil, err
	}
	entity, err := data.ToEntity()
	if err != nil {
		return nil, err
	}
	if err := s.repo.Create(ctx, entity); err != nil {
		return entity, err
	}
	return entity, tx.Commit(ctx)
}

func (s *TabService) CreateManyUserTabs(ctx context.Context, userID uint, data []*tab.CreateDTO) ([]*tab.Tab, error) {
	tx, err := composables.UsePoolTx(ctx)
	if err != nil {
		return nil, err
	}
	entities := make([]*tab.Tab, 0, len(data))
	for _, d := range data {
		entity, err := d.ToEntity()
		if err != nil {
			return nil, err
		}
		entities = append(entities, entity)
	}
	if err := s.repo.CreateMany(ctx, entities); err != nil {
		return nil, err
	}
	if err := s.repo.DeleteUserTabs(ctx, userID); err != nil {
		return nil, err
	}
	return entities, tx.Commit(ctx)
}

func (s *TabService) Update(ctx context.Context, id uint, data *tab.UpdateDTO) error {
	entity, err := data.ToEntity(id)
	if err != nil {
		return err
	}
	if err := s.repo.Update(ctx, entity); err != nil {
		return err
	}
	return nil
}

func (s *TabService) Delete(ctx context.Context, id uint) error {
	return s.repo.Delete(ctx, id)
}
