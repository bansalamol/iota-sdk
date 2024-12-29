package persistence

import (
	"github.com/iota-uz/iota-sdk/modules/core/domain/aggregates/project"
	"github.com/iota-uz/iota-sdk/modules/core/domain/aggregates/role"
	"github.com/iota-uz/iota-sdk/modules/core/domain/aggregates/user"
	"github.com/iota-uz/iota-sdk/modules/core/domain/entities/currency"
	"github.com/iota-uz/iota-sdk/modules/core/domain/entities/employee"
	"github.com/iota-uz/iota-sdk/modules/core/domain/entities/permission"
	stage "github.com/iota-uz/iota-sdk/modules/core/domain/entities/project_stages"
	"github.com/iota-uz/iota-sdk/modules/core/domain/entities/tab"
	"github.com/iota-uz/iota-sdk/modules/core/domain/entities/upload"
	"github.com/iota-uz/iota-sdk/modules/core/infrastructure/persistence/models"
	"time"

	"github.com/iota-agency/iota-sdk/modules/core/infrastructure/persistence/models"

	"github.com/gabriel-vasile/mimetype"
)

func ToDomainUser(dbUser *models.User) (*user.User, error) {
	roles := make([]*role.Role, len(dbUser.Roles))
	// for i, r := range dbUser.Roles {
	// 	roles[i] = toDomainRole(&r)
	// }
	var middleName string
	if dbUser.MiddleName != nil {
		middleName = *dbUser.MiddleName
	}
	var password string
	if dbUser.Password != nil {
		password = *dbUser.Password
	}
	var avatar upload.Upload
	if dbUser.Avatar != nil {
		avatar = *ToDomainUpload(dbUser.Avatar)
	}

	return &user.User{
		ID:         dbUser.ID,
		FirstName:  dbUser.FirstName,
		LastName:   dbUser.LastName,
		MiddleName: middleName,
		Email:      dbUser.Email,
		Password:   password,
		AvatarID:   dbUser.AvatarID,
		Avatar:     &avatar,
		EmployeeID: dbUser.EmployeeID,
		UILanguage: user.UILanguage(dbUser.UiLanguage),
		LastIP:     dbUser.LastIP,
		LastLogin:  dbUser.LastLogin,
		LastAction: dbUser.LastAction,
		CreatedAt:  dbUser.CreatedAt,
		UpdatedAt:  dbUser.UpdatedAt,
		Roles:      roles,
	}, nil
}

func toDBUser(entity *user.User) (*models.User, []models.Role) {
	roles := make([]models.Role, len(entity.Roles))
	for i, r := range entity.Roles {
		dbRole, _ := toDBRole(r)
		roles[i] = *dbRole
	}
	var avatar *models.Upload
	if v := entity.AvatarID; v != nil {
		avatar = ToDBUpload(&upload.Upload{
			ID: *v,
		})
	}
	var avatarId *uint
	if entity.AvatarID != nil && *entity.AvatarID != 0 {
		avatarId = entity.AvatarID
	}
	return &models.User{
		ID:         entity.ID,
		FirstName:  entity.FirstName,
		LastName:   entity.LastName,
		MiddleName: &entity.MiddleName,
		Email:      entity.Email,
		UiLanguage: string(entity.UILanguage),
		Password:   &entity.Password,
		AvatarID:   avatarId,
		EmployeeID: entity.EmployeeID,
		Avatar:     avatar,
		LastIP:     entity.LastIP,
		LastLogin:  entity.LastLogin,
		LastAction: entity.LastAction,
		CreatedAt:  entity.CreatedAt,
		UpdatedAt:  entity.UpdatedAt,
		Roles:      nil,
	}, roles
}

func toDomainRole(dbRole *models.Role) (*role.Role, error) {
	// permissions := make([]permission.Permission, len(dbRole.Permissions))
	// for i, p := range dbRole.Permissions {
	// 	permissions[i] = *toDomainPermission(&p)
	// }
	return &role.Role{
		ID:          dbRole.ID,
		Name:        dbRole.Name,
		Description: dbRole.Description,
		Permissions: make([]permission.Permission, 0),
		// Permissions: permissions,
		CreatedAt: dbRole.CreatedAt,
		UpdatedAt: dbRole.UpdatedAt,
	}, nil
}

func toDBRole(entity *role.Role) (*models.Role, []models.Permission) {
	permissions := make([]models.Permission, len(entity.Permissions))
	for i, p := range entity.Permissions {
		permissions[i] = toDBPermission(p)
	}
	return &models.Role{
		ID:          entity.ID,
		Name:        entity.Name,
		Description: entity.Description,
		Permissions: nil,
		CreatedAt:   entity.CreatedAt,
		UpdatedAt:   entity.UpdatedAt,
	}, permissions
}

func toDBPermission(entity permission.Permission) models.Permission {
	return models.Permission{
		ID:       entity.ID,
		Name:     entity.Name,
		Resource: string(entity.Resource),
		Action:   string(entity.Action),
		Modifier: string(entity.Modifier),
	}
}

func toDomainPermission(dbPermission models.Permission) (permission.Permission, error) {
	return permission.Permission{
		ID:       dbPermission.ID,
		Name:     dbPermission.Name,
		Resource: permission.Resource(dbPermission.Resource),
		Action:   permission.Action(dbPermission.Action),
		Modifier: permission.Modifier(dbPermission.Modifier),
	}, nil
}

func toDomainProject(dbProject *models.Project) *project.Project {
	return &project.Project{
		ID:          dbProject.ID,
		Name:        dbProject.Name,
		Description: dbProject.Description,
		CreatedAt:   dbProject.CreatedAt,
		UpdatedAt:   dbProject.UpdatedAt,
	}
}

func toDBProject(entity *project.Project) *models.Project {
	return &models.Project{
		ID:          entity.ID,
		Name:        entity.Name,
		Description: entity.Description,
		CreatedAt:   entity.CreatedAt,
		UpdatedAt:   entity.UpdatedAt,
	}
}

func toDomainEmployee(dbEmployee *models.Employee) *employee.Employee {
	return &employee.Employee{
		ID:        dbEmployee.ID,
		FirstName: dbEmployee.FirstName,
		LastName:  dbEmployee.LastName,
		Email:     dbEmployee.Email,
		Phone:     dbEmployee.Phone,
		CreatedAt: dbEmployee.CreatedAt,
		UpdatedAt: dbEmployee.UpdatedAt,
		Meta: &employee.Meta{
			EmployeeID: dbEmployee.ID,
		},
	}
}

func toDBEmployee(entity *employee.Employee) (*models.Employee, *models.EmployeeMeta) {
	dbEmployee := &models.Employee{
		ID:        entity.ID,
		FirstName: entity.FirstName,
		LastName:  entity.LastName,
		Email:     entity.Email,
		Phone:     entity.Phone,
		CreatedAt: entity.CreatedAt,
		UpdatedAt: entity.UpdatedAt,
	}
	dbEmployeeMeta := &models.EmployeeMeta{
		EmployeeID: entity.ID,
		UpdatedAt:  entity.UpdatedAt,
	}
	return dbEmployee, dbEmployeeMeta
}

func toDomainProjectStage(dbStage *models.ProjectStage) *stage.ProjectStage {
	return &stage.ProjectStage{
		ID:        dbStage.ID,
		Name:      dbStage.Name,
		CreatedAt: dbStage.CreatedAt,
		UpdatedAt: dbStage.UpdatedAt,
	}
}

func toDBProjectStage(entity *stage.ProjectStage) *models.ProjectStage {
	return &models.ProjectStage{
		ID:        entity.ID,
		Name:      entity.Name,
		ProjectID: entity.ProjectID,
		CreatedAt: entity.CreatedAt,
		UpdatedAt: entity.UpdatedAt,
	}
}

func ToDBUpload(upload *upload.Upload) *models.Upload {
	return &models.Upload{
		ID:        upload.ID,
		Path:      upload.Path,
		Hash:      upload.Hash,
		Size:      upload.Size,
		Mimetype:  upload.Mimetype.String(),
		CreatedAt: upload.CreatedAt,
		UpdatedAt: upload.UpdatedAt,
	}
}

func ToDomainUpload(dbUpload *models.Upload) *upload.Upload {
	var mime mimetype.MIME
	if dbUpload.Mimetype != "" {
		mime = *mimetype.Lookup(dbUpload.Mimetype)
	}
	return &upload.Upload{
		ID:        dbUpload.ID,
		Size:      dbUpload.Size,
		Path:      dbUpload.Path,
		Mimetype:  mime,
		CreatedAt: dbUpload.CreatedAt,
		UpdatedAt: dbUpload.UpdatedAt,
	}
}

func ToDBCurrency(entity *currency.Currency) *models.Currency {
	return &models.Currency{
		Code:      string(entity.Code),
		Name:      entity.Name,
		Symbol:    string(entity.Symbol),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

func ToDomainCurrency(dbCurrency *models.Currency) (*currency.Currency, error) {
	code, err := currency.NewCode(dbCurrency.Code)
	if err != nil {
		return nil, err
	}
	symbol, err := currency.NewSymbol(dbCurrency.Symbol)
	if err != nil {
		return nil, err
	}
	return &currency.Currency{
		Code:   code,
		Name:   dbCurrency.Name,
		Symbol: symbol,
	}, nil
}

func ToDBTab(tab *tab.Tab) *models.Tab {
	return &models.Tab{
		ID:       tab.ID,
		Href:     tab.Href,
		Position: tab.Position,
		UserID:   tab.UserID,
	}
}

func ToDomainTab(dbTab *models.Tab) (*tab.Tab, error) {
	return &tab.Tab{
		ID:       dbTab.ID,
		Href:     dbTab.Href,
		Position: dbTab.Position,
		UserID:   dbTab.UserID,
	}, nil
}

func toDomainPosition(dbPosition *models.Position) (*position.Position, error) {
	return &position.Position{
		ID:          dbPosition.ID,
		Name:        dbPosition.Name,
		Description: dbPosition.Description,
		CreatedAt:   dbPosition.CreatedAt,
		UpdatedAt:   dbPosition.UpdatedAt,
	}, nil
}

func toDBPosition(position *position.Position) *models.Position {
	return &models.Position{
		ID:          position.ID,
		Name:        position.Name,
		Description: position.Description,
		CreatedAt:   position.CreatedAt,
		UpdatedAt:   position.UpdatedAt,
	}
}

func toDBSession(session *session.Session) *models.Session {
	return &models.Session{
		UserID:    session.UserID,
		Token:     session.Token,
		IP:        session.IP,
		UserAgent: session.UserAgent,
		CreatedAt: session.CreatedAt,
		ExpiresAt: session.ExpiresAt,
	}
}

func toDomainSession(dbSession *models.Session) *session.Session {
	return &session.Session{
		UserID:    dbSession.UserID,
		Token:     dbSession.Token,
		IP:        dbSession.IP,
		UserAgent: dbSession.UserAgent,
		CreatedAt: dbSession.CreatedAt,
		ExpiresAt: dbSession.ExpiresAt,
	}
}

func toDBAuthenticationLog(log *authlog.AuthenticationLog) *models.AuthenticationLog {
	return &models.AuthenticationLog{
		ID:        log.ID,
		UserID:    log.UserID,
		IP:        log.IP,
		UserAgent: log.UserAgent,
		CreatedAt: log.CreatedAt,
	}
}

func toDomainAuthenticationLog(dbLog *models.AuthenticationLog) *authlog.AuthenticationLog {
	return &authlog.AuthenticationLog{
		ID:        dbLog.ID,
		UserID:    dbLog.UserID,
		IP:        dbLog.IP,
		UserAgent: dbLog.UserAgent,
		CreatedAt: dbLog.CreatedAt,
	}
}
