package services

import (
	"context"

	prompt2 "github.com/iota-uz/iota-sdk/modules/bichat/domain/entities/prompt"
	"github.com/iota-uz/iota-sdk/pkg/eventbus"
)

type PromptService struct {
	repo      prompt2.Repository
	publisher eventbus.EventBus
}

func NewPromptService(repo prompt2.Repository, publisher eventbus.EventBus) *PromptService {
	return &PromptService{
		repo:      repo,
		publisher: publisher,
	}
}

func (s *PromptService) Count(ctx context.Context) (int64, error) {
	return s.repo.Count(ctx)
}

func (s *PromptService) GetAll(ctx context.Context) ([]*prompt2.Prompt, error) {
	return s.repo.GetAll(ctx)
}

func (s *PromptService) GetByID(ctx context.Context, id string) (*prompt2.Prompt, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *PromptService) GetPaginated(
	ctx context.Context,
	limit, offset int,
	sortBy []string,
) ([]*prompt2.Prompt, error) {
	return s.repo.GetPaginated(ctx, limit, offset, sortBy)
}

func (s *PromptService) Create(ctx context.Context, data *prompt2.Prompt) error {
	if err := s.repo.Create(ctx, data); err != nil {
		return err
	}
	s.publisher.Publish("prompt.created", data)
	return nil
}

func (s *PromptService) Update(ctx context.Context, data *prompt2.Prompt) error {
	if err := s.repo.Update(ctx, data); err != nil {
		return err
	}
	s.publisher.Publish("prompt.updated", data)
	return nil
}

func (s *PromptService) Delete(ctx context.Context, id int64) error {
	if err := s.repo.Delete(ctx, id); err != nil {
		return err
	}
	s.publisher.Publish("prompt.deleted", id)
	return nil
}
