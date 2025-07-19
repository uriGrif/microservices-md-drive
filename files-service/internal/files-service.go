package internal

import (
	"database/sql"
	"log"
)

type FilesService struct {
	db *sql.DB
}

func NewFilesService(db *sql.DB) FilesService {
	if db == nil {
		log.Fatal("failed to create files service: db handle is nil")
	}
	return FilesService{
		db: db,
	}
}

func (fs *FilesService) CreateFile(userId string, name string) (string, error) {
	// create file and owner permissions over it
	// TODO
	return "", nil
}

func (fs *FilesService) GetFile(fileId string, userId string) (File, error) {
	// check permissions and return file
	// TODO
	return File{}, nil
}

func (fs *FilesService) ListFilesByUser(userId string) ([]File, error) {
	// TODO
	return nil, nil
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
