package internal

import (
	"log"

	"github.com/gin-gonic/gin"
)

type FilesController struct {
	service *FilesService
}

func NewFilesController(service *FilesService) *FilesController {
	if service == nil {
		log.Fatal("failed to create files controller: service is nil")
	}
	return &FilesController{
		service: service,
	}
}

func (fc *FilesController) CreateFile(ctx *gin.Context) {
	// TODO
}

func (fs *FilesController) GetFile(ctx *gin.Context) {
	// TODO
}

func (fs *FilesController) ListFilesByUser(ctx *gin.Context) {
	// TODO
}

// update file metadata, not the content, that is managed by the files service
func (fs *FilesController) UpdateFile(ctx *gin.Context) {
	// TODO
}

func (fs *FilesController) DeleteFile(ctx *gin.Context) {
	// TODO
}
