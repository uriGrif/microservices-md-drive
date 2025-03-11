package internal

import (
	"context"
	"database/sql"
	"errors"
	"log"
	"permission-service/protogen/golang/permissions"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

// Note: I am not using a dedicated DB Acces Layer (like a repository), in order to keep the code simple and short in one single file
// Of course, this is not a good practice, because it leads to a lot of duplicated code and makes the code harder to maintain and test
// The main objective of this code is to dive into implementing the gRPC service, so I am taking this shortcut here

type PermissionService struct {
	db *sql.DB
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

func getAuthenticatedUser(ctx context.Context) (string, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return "", errors.New("failed to get metadata from context")
	}
	userID := md.Get("x-authenticated-user")
	if len(userID) == 0 {
		return "", errors.New("user not authenticated")
	}
	return userID[0], nil
}

func (ps *PermissionService) CreatePermission(ctx context.Context, req *permissions.CreatePermissionRequest) (*permissions.CreatePermissionResponse, error) {

	userID, authErr := getAuthenticatedUser(ctx)
	if authErr != nil {
		return nil, status.Errorf(codes.Unauthenticated, authErr.Error())
	}

	if userID != req.Permission.UserId {
		// owner wants to grant permissions to a different user

		// check that the user is the owner of the file
		q := "SELECT 1 FROM permission WHERE file_id = $1 AND user_id = $2 AND permission_level = $3"
		var exists int
		err := ps.db.QueryRow(q, req.Permission.FileId, userID, permissions.PermissionLevel_OWNER).Scan(&exists)
		switch {
		case err == sql.ErrNoRows:
			return nil, status.Errorf(codes.PermissionDenied, "user is not the owner of the file")
		case err != nil:
			log.Printf("query error: %v\n", err)
			return nil, status.Error(codes.Internal, "Internal Server Error")
		}
	} else {
		// owner wants to create a permissions for themselves

		if req.Permission.Level != permissions.PermissionLevel_OWNER {
			// a user can only create an OWNER permission for themselves
			return nil, status.Errorf(codes.PermissionDenied, "user can only create OWNER permission for themselves")
		}

		// check that the file does not have an owner already
		q := "SELECT 1 FROM permission WHERE file_id = $1 AND user_id = $2 AND permission_level = $3"
		var exists int
		err := ps.db.QueryRow(q, req.Permission.FileId, userID, permissions.PermissionLevel_OWNER).Scan(&exists)
		switch {
		case err != sql.ErrNoRows:
			return nil, status.Errorf(codes.PermissionDenied, "file already has an owner")
		case err != nil:
			log.Printf("query error: %v\n", err)
			return nil, status.Error(codes.Internal, "Internal Server Error")
		}
	}

	q := "INSERT INTO permission(file_id, user_id, permission_level) values ($1, $2, $3)"
	_, err := ps.db.Exec(q, req.Permission.FileId, req.Permission.UserId, req.Permission.Level)
	if err != nil {
		return nil, err
	}
	return &permissions.CreatePermissionResponse{
		Permission: req.Permission,
	}, nil
}

func (ps *PermissionService) GetPermission(ctx context.Context, req *permissions.GetPermissionRequest) (*permissions.GetPermissionResponse, error) {
	q := "SELECT file_id, user_id, permission_level FROM permission WHERE file_id = $1 AND user_id = $2"
	var permission permissions.Permission
	err := ps.db.QueryRow(q, req.FileId, req.UserId).Scan(permission.FileId, permission.UserId, permission.Level)
	switch {
	case err != sql.ErrNoRows:
		return nil, status.Errorf(codes.NotFound, "permission not found")
	case err != nil:
		log.Printf("query error: %v\n", err)
		return nil, status.Error(codes.Internal, "Internal Server Error")
	}
	return &permissions.GetPermissionResponse{
		Permission: &permission,
	}, nil
}

func (ps *PermissionService) ListPermissionsByFile(ctx context.Context, req *permissions.ListPermissionsByFileRequest) (*permissions.ListPermissionsByFileResponse, error) {
	q := "SELECT file_id, user_id, permission_level FROM permission WHERE file_id = $1"
	var results []*permissions.Permission
	rows, err := ps.db.Query(q, req.FileId)
	if err != nil {
		log.Printf("query error: %v\n", err)
		return nil, status.Error(codes.Internal, "Internal Server Error")
	}
	for rows.Next() {
		var p permissions.Permission
		err := rows.Scan(p.FileId, p.UserId, p.Level)
		if err != nil {
			log.Printf("query error: %v\n", err)
			return nil, status.Error(codes.Internal, "Internal Server Error")
		}
		results = append(results, &p)
	}
	return &permissions.ListPermissionsByFileResponse{
		Permissions: results,
	}, nil
}

func (ps *PermissionService) ListPermissionsByUser(ctx context.Context, req *permissions.ListPermissionsByUserRequest) (*permissions.ListPermissionsByUserResponse, error) {
	userID, authErr := getAuthenticatedUser(ctx)
	if authErr != nil {
		return nil, status.Errorf(codes.Unauthenticated, authErr.Error())
	}
	if userID != req.UserId {
		return nil, status.Errorf(codes.PermissionDenied, "user can only list their own permissions")
	}

	q := "SELECT file_id, user_id, permission_level FROM permission WHERE user_id = $1"
	var results []*permissions.Permission
	rows, err := ps.db.Query(q, req.UserId)
	if err != nil {
		log.Printf("query error: %v\n", err)
		return nil, status.Error(codes.Internal, "Internal Server Error")
	}
	for rows.Next() {
		var p permissions.Permission
		err := rows.Scan(p.FileId, p.UserId, p.Level)
		if err != nil {
			log.Printf("query error: %v\n", err)
			return nil, status.Error(codes.Internal, "Internal Server Error")
		}
		results = append(results, &p)
	}
	return &permissions.ListPermissionsByUserResponse{
		Permissions: results,
	}, nil
}

func (ps *PermissionService) UpdatePermission(ctx context.Context, req *permissions.UpdatePermissionRequest) (*permissions.UpdatePermissionResponse, error) {
	userID, authErr := getAuthenticatedUser(ctx)
	if authErr != nil {
		return nil, status.Errorf(codes.Unauthenticated, authErr.Error())
	}

	// check that user is the owner of the file
	q := "SELECT 1 FROM permission WHERE file_id = $1 AND user_id = $2 AND permission_level = $3"
	var exists int
	err := ps.db.QueryRow(q, req.Permission.FileId, userID, permissions.PermissionLevel_OWNER).Scan(&exists)
	switch {
	case err == sql.ErrNoRows:
		return nil, status.Errorf(codes.PermissionDenied, "user cannot modify permissions to the file, as they are not the owner")
	case err != nil:
		log.Printf("query error: %v\n", err)
		return nil, status.Error(codes.Internal, "Internal Server Error")
	}

	q = "UPDATE permission SET permission_level = $1 WHERE file_id = $2 AND user_id = $3"
	_, err = ps.db.Exec(q, req.Permission.Level, req.Permission.FileId, req.Permission.UserId)
	if err != nil {
		return nil, err
	}
	return &permissions.UpdatePermissionResponse{
		Permission: req.Permission,
	}, nil
}

func (ps *PermissionService) DeletePermission(ctx context.Context, req *permissions.DeletePermissionRequest) (*permissions.DeletePermissionResponse, error) {
	userID, authErr := getAuthenticatedUser(ctx)
	if authErr != nil {
		return nil, status.Errorf(codes.Unauthenticated, authErr.Error())
	}

	// check that user is the owner of the file
	q := "SELECT 1 FROM permission WHERE file_id = $1 AND user_id = $2 AND permission_level = $3"
	var exists int
	err := ps.db.QueryRow(q, req.FileId, userID, permissions.PermissionLevel_OWNER).Scan(&exists)
	switch {
	case err == sql.ErrNoRows:
		return nil, status.Errorf(codes.PermissionDenied, "user cannot remove permissions to the file, as they are not the owner")
	case err != nil:
		log.Printf("query error: %v\n", err)
		return nil, status.Error(codes.Internal, "Internal Server Error")
	}

	q = "DELETE FROM permission WHERE WHERE file_id = $1 AND user_id = $2"
	_, err = ps.db.Exec(q, req.FileId, req.UserId)
	if err != nil {
		return nil, err
	}
	return &permissions.DeletePermissionResponse{}, nil
}
