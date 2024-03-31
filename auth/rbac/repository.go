package rbac

import (
	"pro_pay/tools/logger"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
)

type Repository struct {
	db     *sqlx.DB
	logrus *logger.Logger
}

func NewRepository(db *sqlx.DB, logrus *logger.Logger) *Repository {
	return &Repository{db: db, logrus: logrus}
}

func (r *Repository) GetPermissions() (permissions []Permissions, err error) {
	query := `SELECT id, router ,method , (select ARRAY_AGG(role_id)::uuid[] from role_permission rp where p.id =rp.permission_id) role_ids FROM "public"."permission" p  WHERE p.deleted_at IS NULL`
	row, err := r.db.Query(query)
	if err != nil {
		r.logrus.Error(err)
		return permissions, err
	}
	for row.Next() {
		var roles []uuid.UUID
		permission := Permissions{}
		err = row.Scan(
			&permission.ID,
			&permission.Router,
			&permission.Method,
			pq.Array(&roles),
		)
		if err != nil {
			r.logrus.Error(err)
			return permissions, err
		}
		permission.RoleIDs = roles
		permissions = append(permissions, permission)
	}
	return permissions, nil
}
