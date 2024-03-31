package rbac

import (
	"github.com/google/uuid"
)

type Permissions struct {
	ID      uuid.UUID   `db:"id"`
	Router  string      `db:"router"`
	Method  string      `db:"method"`
	RoleIDs []uuid.UUID `db:"role_ids"`
}
