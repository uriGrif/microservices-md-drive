package internal

import (
	"database/sql"
	"permission-service/protogen/golang/permissions"
)

type PermissionService struct {
	db *sql.DB // It would be better to have an extra DB Access layer, like a repository, but I won't do it here to keep the code simpler and shorter
	permissions.UnimplementedPermissionServiceServer
}

func NewPermissionService(db *sql.DB) PermissionService {
	return PermissionService{
		db: db,
	}
}

// func (ps *PermissionService) CreatePermission(context.Context, *permissions.CreatePermissionRequest) (*permissions.CreatePermissionResponse, error) {
// }

// func (ps *PermissionService) GetPermission(context.Context, *permissions.GetPermissionRequest) (*permissions.GetPermissionResponse, error) {
// }

// func (ps *PermissionService) ListPermissions(context.Context, *permissions.ListPermissionsRequest) (*permissions.ListPermissionsResponse, error) {
// }

// func (ps *PermissionService) UpdatePermission(context.Context, *permissions.UpdatePermissionRequest) (*permissions.UpdatePermissionResponse, error) {
// }

// func (ps *PermissionService) DeletePermission(context.Context, *permissions.DeletePermissionRequest) (*permissions.DeletePermissionResponse, error) {
// }
