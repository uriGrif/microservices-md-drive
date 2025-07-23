package internal

import (
	"context"
	"database/sql"
	permissions "files-service/protogen/golang"
	"fmt"
	"log"
	"strings"

	"github.com/google/uuid"
	"google.golang.org/grpc/metadata"
)

type FilesService struct {
	db                *sql.DB
	permissionsClient permissions.PermissionServiceClient
}

func NewFilesService(db *sql.DB, permissionsClient permissions.PermissionServiceClient) FilesService {
	if db == nil {
		log.Fatal("failed to create files service: db handle is nil")
	}
	if permissionsClient == nil {
		log.Fatal("failed to create files service: permissions client is nil")
	}
	return FilesService{
		db:                db,
		permissionsClient: permissionsClient,
	}
}

func generateFileId() string {
	return uuid.NewString()
}

func (fs *FilesService) CreateFile(userId string, name string) (*File, error) {
	// create file and owner permissions over it
	query := "INSERT INTO files (id, user_id, name) VALUES (?, ?, ?)"
	fileId := generateFileId()
	_, err := fs.db.Exec(query, fileId, userId, name)
	if err != nil {
		return nil, err
	}

	md := metadata.Pairs("x-authenticated-user", userId)
	ctxWithMetadata := metadata.NewOutgoingContext(context.Background(), md)
	_, permissionErr := fs.permissionsClient.CreatePermission(ctxWithMetadata, &permissions.CreatePermissionRequest{
		Permission: &permissions.Permission{
			FileId: fileId,
			UserId: userId,
			Level:  permissions.PermissionLevel_OWNER,
		},
	})
	if permissionErr != nil {
		log.Printf("failed to create owner permission for file %s: %v", fileId, permissionErr)
		return nil, permissionErr
	}
	return &File{
		Id:     fileId,
		UserId: userId,
		Name:   name,
	}, nil
}

func (fs *FilesService) GetFile(fileId string, userId string) (File, error) {
	query := "SELECT id, user_id, name, created_at, updated_at FROM files WHERE id = ?"

	md := metadata.Pairs("x-authenticated-user", userId)
	ctxWithMetadata := metadata.NewOutgoingContext(context.Background(), md)
	_, err := fs.permissionsClient.GetPermission(ctxWithMetadata, &permissions.GetPermissionRequest{
		FileId: fileId,
		UserId: userId,
	})
	if err != nil {
		log.Printf("failed to get permission for file %s: %v", fileId, err)
		return File{}, err
	}
	row := fs.db.QueryRow(query, fileId)
	var file File
	err = row.Scan(&file.Id, &file.UserId, &file.Name, &file.CreatedAt, &file.UpdatedAt)
	if err != nil {
		return File{}, err
	}
	return file, nil
}

func (fs *FilesService) ListFilesByUser(userId string) ([]ListFile, error) {
	md := metadata.Pairs("x-authenticated-user", userId)
	ctxWithMetadata := metadata.NewOutgoingContext(context.Background(), md)
	permissionsRes, err := fs.permissionsClient.ListPermissionsByUser(ctxWithMetadata, &permissions.ListPermissionsByUserRequest{
		UserId: userId,
	})
	if err != nil {
		log.Printf("failed to get permissions for user %s: %v", userId, err)
		return nil, err
	}
	if len(permissionsRes.Permissions) == 0 {
		return []ListFile{}, nil // no files found for user
	}

	fileIds := make([]interface{}, 0, len(permissionsRes.Permissions))
	permissionsMap := make(map[string]permissions.PermissionLevel)
	for _, p := range permissionsRes.Permissions {
		fileIds = append(fileIds, p.FileId)
		permissionsMap[p.FileId] = p.Level
	}

	query := fmt.Sprintf("SELECT id, user_id, name, created_at FROM files WHERE id IN (%s)", strings.Repeat("?,", len(fileIds)-1)+"?")
	rows, err := fs.db.Query(query, fileIds...)
	if err != nil {
		log.Printf("failed to list files for user %s: %v", userId, err)
		return nil, err
	}

	var result []ListFile

	defer rows.Close()
	for rows.Next() {
		var file ListFile
		err := rows.Scan(&file.Id, &file.OwnerId, &file.Name, &file.CreatedAt)
		if err != nil {
			log.Printf("failed to scan file row: %v", err)
			return nil, err
		}
		file.PermissionLevel = permissionsMap[file.Id].String()
		result = append(result, file)
	}

	return result, nil
}

// update file metadata, not the content, that is managed by the files service
func (fs *FilesService) UpdateFile(fileId string, userId string, file File) (File, error) {
	// TODO
	return file, nil
}

func (fs *FilesService) DeleteFile(fileId string, userId string) error {
	// TODO
	return nil
}
