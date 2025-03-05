package internal

import (
	"context"
	"database/sql"
	"log"
	"permission-service/protogen/golang/permissions"
)

type PermissionService struct {
	db *sql.DB // It would be better to have an extra DB Access layer, like a repository, but I won't do it here to keep the code simpler and shorter
	permissions.UnimplementedPermissionServiceServer
}

func NewPermissionService(db *sql.DB) PermissionService {
	if db == nil {
		log.Fatal("failed to create permission service: db handle is nil")
	}
	return PermissionService{
		db: db,
	}
}

func (ps *PermissionService) CreatePermission(ctx context.Context, p *permissions.CreatePermissionRequest) (*permissions.CreatePermissionResponse, error) {
	q := "INSERT INTO permission(file_id, user_id, permission_level) values ($1, $2, $3)"
	_, err := ps.db.Exec(q, p.Permission.FileId, p.Permission.UserId, p.Permission.Level)
	if err != nil {
		return nil, err
	}
	return &permissions.CreatePermissionResponse{
		Permission: p.Permission,
	}, nil
}

// func (ps *PermissionService) GetPermission(context.Context, *permissions.GetPermissionRequest) (*permissions.GetPermissionResponse, error) {
// }

// func (ps *PermissionService) ListPermissions(context.Context, *permissions.ListPermissionsRequest) (*permissions.ListPermissionsResponse, error) {
// }

// func (ps *PermissionService) UpdatePermission(context.Context, *permissions.UpdatePermissionRequest) (*permissions.UpdatePermissionResponse, error) {
// }

// func (ps *PermissionService) DeletePermission(context.Context, *permissions.DeletePermissionRequest) (*permissions.DeletePermissionResponse, error) {
// }
