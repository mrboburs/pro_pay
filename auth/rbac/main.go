package rbac

import (
	"pro_pay/tools/logger"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

var rbac *RBAC

type RBAC struct {
	Permissions          map[string]map[uuid.UUID]struct{}
	PermissionsDB        []Permissions
	PermissionsHashTable map[uuid.UUID]Permissions
	logger               *logger.Logger
	Methods
}

type Methods interface {
	GetPermissions() ([]Permissions, error)
}

func NewRBAC(db *sqlx.DB, logger *logger.Logger) *RBAC {
	rbac = &RBAC{
		Methods:              NewRepository(db, logger),
		Permissions:          make(map[string]map[uuid.UUID]struct{}),
		PermissionsDB:        []Permissions{},
		PermissionsHashTable: make(map[uuid.UUID]Permissions),
		logger:               logger,
	}
	StartRBAC(logger, rbac)
	return rbac
}

func StartRBAC(logger *logger.Logger, rbac *RBAC) {
	permissions, err := rbac.GetPermissions()
	if err != nil {
		logger.Error(err)
		return
	}
	rbac.PermissionsDB = permissions
	rbac.PermissionsConvert()
}
