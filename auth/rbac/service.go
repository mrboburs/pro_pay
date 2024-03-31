package rbac

import (
	"errors"

	"github.com/google/uuid"
)

var (
	errNotFound         = errors.New("not found resource")
	errPermissionDenied = errors.New("permission denied")
)

func (rbac *RBAC) PermissionsConvert() {
	permissions := rbac.PermissionsDB
	for _, p := range permissions {
		resource := GetResource(p.Router, p.Method)
		_, resourceExist := rbac.Permissions[resource]
		if !resourceExist {
			rolesMap := make(map[uuid.UUID]struct{})
			for _, r := range p.RoleIDs {
				rolesMap[r] = struct{}{}
			}
			rbac.Permissions[resource] = rolesMap
		}
		_, permissionExist := rbac.PermissionsHashTable[p.ID]
		if !permissionExist {
			rbac.PermissionsHashTable[p.ID] = p
		}
	}
}

func HasPermission(roleID uuid.UUID, router, method string) error {
	resource := GetResource(router, method)
	_, ok := rbac.Permissions[resource]
	if !ok {
		rbac.logger.Error(errNotFound)
		return errNotFound
	}
	_, ok = rbac.Permissions[resource][roleID]
	if !ok {
		rbac.logger.Error(errPermissionDenied)
		return errPermissionDenied
	}
	return nil
}

func AddPermission(permissionID uuid.UUID, router, method string) (err error) {
	rbac.PermissionsHashTable[permissionID] = Permissions{
		ID:      permissionID,
		Router:  router,
		Method:  method,
		RoleIDs: []uuid.UUID{},
	}
	resource := GetResource(router, method)
	rbac.Permissions[resource] = make(map[uuid.UUID]struct{})
	return nil
}

func UpdatePermissionRole(permissionID, roleID uuid.UUID) (err error) {
	permission, ok := rbac.PermissionsHashTable[permissionID]
	if !ok {
		rbac.logger.Error(errNotFound)
		return errNotFound
	}
	roles := permission.RoleIDs
	roleExist := false
	for i := range permission.RoleIDs {
		if permission.RoleIDs[i] == roleID {
			roles = deleteSliceItem(roles, i)
			roleExist = true
			break
		}
	}
	if !roleExist {
		roles = append(roles, roleID)
	}
	rbac.PermissionsHashTable[permissionID] = Permissions{
		ID:      permissionID,
		Router:  permission.Router,
		Method:  permission.Method,
		RoleIDs: roles,
	}
	resource := GetResource(permission.Router, permission.Method)
	_, ok = rbac.Permissions[resource]
	if !ok {
		rbac.logger.Error(errNotFound)
		return errNotFound
	}
	_, ok = rbac.Permissions[resource][roleID]
	if !ok {
		rbac.Permissions[resource][roleID] = struct{}{}
	} else {
		// for k := range rbac.Permissions[resource] {
		// 	if k == roleID {
		// 		continue
		// 	}
		// 	rbac.Permissions[resource][k] = struct{}{}
		// }
		delete(rbac.Permissions[resource], roleID)
		rbac.logger.Error(rbac.Permissions[resource], roleID)
	}
	return nil
}
func UpdatePermission(permissionID uuid.UUID, router, method string) (err error) {
	permission, ok := rbac.PermissionsHashTable[permissionID]
	if !ok {
		rbac.logger.Error(errNotFound)
		return errNotFound
	}
	resource := GetResource(permission.Router, permission.Method)
	delete(rbac.Permissions, resource)
	rbac.PermissionsHashTable[permissionID] = Permissions{
		ID:      permissionID,
		Router:  router,
		Method:  method,
		RoleIDs: permission.RoleIDs,
	}
	resource = GetResource(router, method)
	rbac.Permissions[resource] = make(map[uuid.UUID]struct{})
	for _, r := range permission.RoleIDs {
		rbac.Permissions[resource][r] = struct{}{}
	}
	return nil
}
func DeletePermission(permissionID uuid.UUID) (err error) {
	permission, ok := rbac.PermissionsHashTable[permissionID]
	if !ok {
		rbac.logger.Error(errNotFound)
		return errNotFound
	}
	resource := GetResource(permission.Router, permission.Method)
	delete(rbac.Permissions, resource)
	delete(rbac.PermissionsHashTable, permissionID)
	return nil
}

func GetResource(router, method string) string {
	resource := method + "|" + router
	return resource
}

func deleteSliceItem(roles []uuid.UUID, index int) []uuid.UUID {
	if index >= 0 && index < len(roles)-1 {
		roles = append(roles[:index], roles[index+1:]...)
	} else if index == 0 {
		roles = roles[1:]
	} else if index == len(roles)-1 {
		roles = roles[:index]
	}
	return roles
}
