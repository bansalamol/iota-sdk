package controllers

import (
	"fmt"
	"net/http"

	"github.com/a-h/templ"
	"github.com/go-faster/errors"
	"github.com/gorilla/mux"
	"github.com/nicksnyder/go-i18n/v2/i18n"

	"github.com/iota-uz/iota-sdk/components/base/pagination"
	"github.com/iota-uz/iota-sdk/modules/crm/domain/aggregates/client"
	"github.com/iota-uz/iota-sdk/modules/crm/presentation/mappers"
	chatsui "github.com/iota-uz/iota-sdk/modules/crm/presentation/templates/pages/chats"
	"github.com/iota-uz/iota-sdk/modules/crm/presentation/templates/pages/clients"
	"github.com/iota-uz/iota-sdk/modules/crm/presentation/viewmodels"
	"github.com/iota-uz/iota-sdk/modules/crm/services"
	"github.com/iota-uz/iota-sdk/pkg/application"
	"github.com/iota-uz/iota-sdk/pkg/composables"
	"github.com/iota-uz/iota-sdk/pkg/mapping"
	"github.com/iota-uz/iota-sdk/pkg/middleware"
	"github.com/iota-uz/iota-sdk/pkg/shared"
)

type ClientController struct {
	app           application.Application
	clientService *services.ClientService
	chatService   *services.ChatService
	basePath      string
}

type ClientsPaginatedResponse struct {
	Clients         []*viewmodels.Client
	PaginationState *pagination.State
}

func NewClientController(app application.Application, basePath string) application.Controller {
	return &ClientController{
		app:           app,
		clientService: app.Service(services.ClientService{}).(*services.ClientService),
		chatService:   app.Service(services.ChatService{}).(*services.ChatService),
		basePath:      basePath,
	}
}

func (c *ClientController) Key() string {
	return c.basePath
}

func (c *ClientController) Register(r *mux.Router) {
	commonMiddleware := []mux.MiddlewareFunc{
		middleware.Authorize(),
		middleware.RedirectNotAuthenticated(),
		middleware.ProvideUser(),
		middleware.WithLocalizer(c.app.Bundle()),
		middleware.WithPageContext(),
	}
	router := r.PathPrefix(c.basePath).Subrouter()
	router.Use(commonMiddleware...)
	router.Use(middleware.Tabs(), middleware.NavItems())
	router.HandleFunc("", c.List).Methods(http.MethodGet)
	router.HandleFunc("/new", c.GetNew).Methods(http.MethodGet)
	router.HandleFunc("", c.Create).Methods(http.MethodPost)
	router.HandleFunc("/{id:[0-9]+}", c.Update).Methods(http.MethodPost)
	router.HandleFunc("/{id:[0-9]+}", c.Delete).Methods(http.MethodDelete)

	hxRouter := r.PathPrefix(c.basePath).Subrouter()
	hxRouter.Use(commonMiddleware...)
	hxRouter.HandleFunc("/{id:[0-9]+}", c.View).Methods(http.MethodGet)
	hxRouter.HandleFunc("/{id:[0-9]+}/edit", c.GetEdit).Methods(http.MethodGet)
	hxRouter.HandleFunc("/{id:[0-9]+}/tab/{tab:[a-z]+}", c.TabContents).Methods(http.MethodGet)
}

func (c *ClientController) viewModelClients(r *http.Request) (*ClientsPaginatedResponse, error) {
	paginationParams := composables.UsePaginated(r)
	params, err := composables.UseQuery(&client.FindParams{
		Limit:  paginationParams.Limit,
		Offset: paginationParams.Offset,
		SortBy: client.SortBy{
			Fields:    []client.Field{client.CreatedAt},
			Ascending: false,
		},
	}, r)
	if err != nil {
		return nil, errors.Wrap(err, "Error using query")
	}

	expenseEntities, err := c.clientService.GetPaginated(r.Context(), params)
	if err != nil {
		return nil, errors.Wrap(err, "Error retrieving expenses")
	}

	total, err := c.clientService.Count(r.Context())
	if err != nil {
		return nil, errors.Wrap(err, "Error counting expenses")
	}

	return &ClientsPaginatedResponse{
		Clients:         mapping.MapViewModels(expenseEntities, mappers.ClientToViewModel),
		PaginationState: pagination.New(c.basePath, paginationParams.Page, int(total), params.Limit),
	}, nil
}
func (c *ClientController) List(w http.ResponseWriter, r *http.Request) {
	paginated, err := c.viewModelClients(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	props := &clients.IndexPageProps{
		NewURL:          fmt.Sprintf("%s/new", c.basePath),
		Clients:         paginated.Clients,
		PaginationState: paginated.PaginationState,
	}
	if shared.IsHxRequest(r) {
		templ.Handler(clients.ClientsTable(props), templ.WithStreaming()).ServeHTTP(w, r)
	} else {
		templ.Handler(clients.Index(props), templ.WithStreaming()).ServeHTTP(w, r)
	}
}

func (c *ClientController) GetNew(w http.ResponseWriter, r *http.Request) {
	props := &clients.CreatePageProps{
		Client:  &viewmodels.Client{},
		SaveURL: c.basePath,
	}
	templ.Handler(clients.New(props), templ.WithStreaming()).ServeHTTP(w, r)
}

func (c *ClientController) Create(w http.ResponseWriter, r *http.Request) {
	dto, err := composables.UseForm(&client.CreateDTO{}, r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if errorsMap, ok := dto.Ok(r.Context()); !ok {
		props := &clients.CreatePageProps{
			Errors: errorsMap,
			Client: &viewmodels.Client{
				FirstName:  dto.FirstName,
				LastName:   dto.LastName,
				MiddleName: dto.MiddleName,
				Phone:      dto.Phone,
			},
			SaveURL: fmt.Sprintf("%s", c.basePath),
		}
		templ.Handler(clients.CreateForm(props), templ.WithStreaming()).ServeHTTP(w, r)
		return
	}

	if _, err := c.clientService.Create(r.Context(), dto); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	shared.Redirect(w, r, c.basePath)
}

func (c *ClientController) GetEdit(w http.ResponseWriter, r *http.Request) {
	id, err := shared.ParseID(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	entity, err := c.clientService.GetByID(r.Context(), id)
	if err != nil {
		http.Error(w, "Error retrieving client", http.StatusInternalServerError)
		return
	}
	props := &clients.EditPageProps{
		Client:    mappers.ClientToViewModel(entity),
		Errors:    map[string]string{},
		SaveURL:   fmt.Sprintf("%s/%d", c.basePath, id),
		DeleteURL: fmt.Sprintf("%s/%d", c.basePath, id),
	}
	viewLayoutProps := c.buildViewLayoutProps(
		composables.MustUseLocalizer(r.Context()),
		entity,
		true,
	)
	ctx := templ.WithChildren(r.Context(), clients.EditForm(props))
	templ.Handler(clients.ViewComponent(viewLayoutProps), templ.WithStreaming()).ServeHTTP(w, r.WithContext(ctx))
}

func (c *ClientController) buildViewLayoutProps(
	localizer *i18n.Localizer,
	entity client.Client,
	isEditing bool,
) *clients.ViewPageProps {
	clientURL := fmt.Sprintf("%s/%d", c.basePath, entity.ID())
	return &clients.ViewPageProps{
		EditURL:   fmt.Sprintf("%s/edit", clientURL),
		ClientURL: clientURL,
		Client:    mappers.ClientToViewModel(entity),
		IsEditing: isEditing,
		Tabs: []clients.ClientTab{
			{
				Name: localizer.MustLocalize(&i18n.LocalizeConfig{
					MessageID: "Clients.Tabs.General",
				}),
				URL: fmt.Sprintf("%s/tab/general", clientURL),
			},
			{
				Name: localizer.MustLocalize(&i18n.LocalizeConfig{
					MessageID: "Clients.Tabs.Chat",
				}),
				URL: fmt.Sprintf("%s/tab/chat", clientURL),
			},
		},
	}
}

func (c *ClientController) tabToComponent(r *http.Request, clientID uint, tab string) (templ.Component, error) {
	switch tab {
	case "":
		fallthrough
	case "general":
		return clients.General(), nil
	case "chat":
		entity, err := c.clientService.GetByID(r.Context(), clientID)
		if err != nil {
			return nil, errors.Wrap(err, "Error retrieving client")
		}
		chatEntity, err := c.chatService.GetByClientID(r.Context(), clientID)
		if err != nil {
			return nil, errors.Wrap(err, "Error retrieving chat")
		}
		return clients.Chats(chatsui.SelectedChatProps{
			Chat:       mappers.ChatToViewModel(chatEntity, entity),
			ClientsURL: c.basePath,
		}), nil
	default:
		return clients.NotFound(), nil
	}
}

func (c *ClientController) View(w http.ResponseWriter, r *http.Request) {
	clientID, err := shared.ParseID(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	entity, err := c.clientService.GetByID(r.Context(), clientID)
	if err != nil {
		http.Error(w, "Error retrieving client", http.StatusInternalServerError)
		return
	}
	tab := r.URL.Query().Get("tab")
	component, err := c.tabToComponent(r, clientID, tab)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	ctx := templ.WithChildren(r.Context(), component)
	props := c.buildViewLayoutProps(
		composables.MustUseLocalizer(r.Context()),
		entity,
		false,
	)
	templ.Handler(clients.ViewComponent(props), templ.WithStreaming()).ServeHTTP(w, r.WithContext(ctx))
}

func (c *ClientController) TabContents(w http.ResponseWriter, r *http.Request) {
	clientID, err := shared.ParseID(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	component, err := c.tabToComponent(r, clientID, mux.Vars(r)["tab"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	templ.Handler(component, templ.WithStreaming()).ServeHTTP(w, r)
}

func (c *ClientController) Update(w http.ResponseWriter, r *http.Request) {
	id, err := shared.ParseID(r)
	if err != nil {
		http.Error(w, "Error parsing id", http.StatusInternalServerError)
		return
	}
	dto, err := composables.UseForm(&client.UpdateDTO{}, r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if errorsMap, ok := dto.Ok(r.Context()); !ok {
		entity, err := c.clientService.GetByID(r.Context(), id)
		if err != nil {
			http.Error(w, "Error retrieving expense", http.StatusInternalServerError)
			return
		}
		props := &clients.EditPageProps{
			Client:    mappers.ClientToViewModel(entity),
			Errors:    errorsMap,
			SaveURL:   fmt.Sprintf("%s/%d", c.basePath, id),
			DeleteURL: fmt.Sprintf("%s/%d", c.basePath, id),
		}
		templ.Handler(clients.EditForm(props), templ.WithStreaming()).ServeHTTP(w, r)
		return
	}
	if err := c.clientService.Update(r.Context(), id, dto); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	shared.Redirect(w, r, c.basePath)
}

func (c *ClientController) Delete(w http.ResponseWriter, r *http.Request) {
	id, err := shared.ParseID(r)
	if err != nil {
		http.Error(w, "Error parsing id", http.StatusInternalServerError)
		return
	}

	if _, err := c.clientService.Delete(r.Context(), id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	shared.Redirect(w, r, c.basePath)
}
