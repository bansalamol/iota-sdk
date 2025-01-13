package user

import (
	"context"
	"fmt"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/iota-uz/iota-sdk/modules/core/domain/aggregates/role"
	"github.com/iota-uz/iota-sdk/pkg/constants"
)

type CreateDTO struct {
	FirstName  string `validate:"required"`
	LastName   string `validate:"required"`
	Email      string `validate:"required,email"`
	Password   string
	RoleIDs    []uint `validate:"required"`
	AvatarID   uint
	UILanguage string `validate:"required"`
}

type UpdateDTO struct {
	FirstName  string `validate:"required"`
	LastName   string `validate:"required"`
	Email      string `validate:"required,email"`
	Password   string
	RoleIDs    []uint
	AvatarID   uint
	UILanguage string
}

func (u *CreateDTO) Ok(ctx context.Context) (map[string]string, bool) {
	// TODO: Use composables.UseLocalizer(ctx) instead of ctx.Value(constants.LocalizerKey)
	l, ok := ctx.Value(constants.LocalizerKey).(*i18n.Localizer)
	if !ok {
		panic("localizer not found in context")
	}
	errorMessages := map[string]string{}
	errs := constants.Validate.Struct(u)
	if errs == nil {
		return errorMessages, true
	}
	for _, err := range errs.(validator.ValidationErrors) {
		translatedFieldName := l.MustLocalize(&i18n.LocalizeConfig{
			MessageID: fmt.Sprintf("Users.Single.%s", err.Field()),
		})
		errorMessages[err.Field()] = l.MustLocalize(&i18n.LocalizeConfig{
			MessageID: fmt.Sprintf("ValidationErrors.%s", err.Tag()),
			TemplateData: map[string]string{
				"Field": translatedFieldName,
			},
		})
	}

	return errorMessages, len(errorMessages) == 0
}

func (u *UpdateDTO) Ok(ctx context.Context) (map[string]string, bool) {
	// TODO: Use composables.UseLocalizer(ctx) instead of ctx.Value(constants.LocalizerKey)
	l, ok := ctx.Value(constants.LocalizerKey).(*i18n.Localizer)
	if !ok {
		panic("localizer not found in context")
	}
	errorMessages := map[string]string{}
	errs := constants.Validate.Struct(u)
	if errs == nil {
		return errorMessages, true
	}
	for _, err := range errs.(validator.ValidationErrors) {
		translatedFieldName := l.MustLocalize(&i18n.LocalizeConfig{
			MessageID: fmt.Sprintf("Users.Single.%s", err.Field()),
		})
		errorMessages[err.Field()] = l.MustLocalize(&i18n.LocalizeConfig{
			MessageID: fmt.Sprintf("ValidationErrors.%s", err.Tag()),
			TemplateData: map[string]string{
				"Field": translatedFieldName,
			},
		})
	}

	return errorMessages, len(errorMessages) == 0
}

func (u *CreateDTO) ToEntity() (User, error) {
	roles := make([]role.Role, len(u.RoleIDs))
	for i, id := range u.RoleIDs {
		r, err := role.NewWithID(id, "", "", nil, time.Now(), time.Now())
		if err != nil {
			return nil, err
		}
		roles[i] = r
	}
	return &user{
		firstName:  u.FirstName,
		lastName:   u.LastName,
		email:      u.Email,
		roles:      roles,
		password:   u.Password,
		lastLogin:  time.Now(),
		lastAction: time.Now(),
		lastIP:     "",
		avatarID:   u.AvatarID,
		uiLanguage: UILanguage(u.UILanguage),
		employeeID: 0,
		createdAt:  time.Now(),
		updatedAt:  time.Now(),
	}, nil
}

func (u *UpdateDTO) ToEntity(id uint) (User, error) {
	roles := make([]role.Role, len(u.RoleIDs))
	for i, rID := range u.RoleIDs {
		r, err := role.NewWithID(rID, "", "", nil, time.Now(), time.Now())
		if err != nil {
			return nil, err
		}
		roles[i] = r
	}
	return &user{
		id:         id,
		firstName:  u.FirstName,
		lastName:   u.LastName,
		email:      u.Email,
		roles:      roles,
		password:   u.Password,
		lastLogin:  time.Now(),
		lastAction: time.Now(),
		lastIP:     "",
		avatarID:   u.AvatarID,
		uiLanguage: UILanguage(u.UILanguage),
		employeeID: 0,
		createdAt:  time.Now(),
		updatedAt:  time.Now(),
	}, nil
}
