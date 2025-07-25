package internal

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type FilesController struct {
	service *FilesService
}

func authenticateUser(ctx *gin.Context) (string, error) {
	userId := ctx.GetHeader("x-authenticated-user")
	if userId == "" {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "user is not authenticated"})
		return "", errors.New("user is not authenticated")
	}
	return userId, nil
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
	userId, authErr := authenticateUser(ctx)
	if authErr != nil {
		return
	}
	var requestBody CreateFileRequest
	if err := ctx.ShouldBindJSON(&requestBody); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}
	fmt.Printf("Creating file for user %s with name %s\n", userId, requestBody.Name)
	file, err := fc.service.CreateFile(userId, requestBody.Name)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create file"})
		return
	}
	ctx.JSON(http.StatusCreated, file)
}

func (fs *FilesController) GetFile(ctx *gin.Context) {
	userId, authErr := authenticateUser(ctx)
	if authErr != nil {
		return
	}
	file, err := fs.service.GetFile(ctx.Param("id"), userId)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "file not found"})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get file"})
		return
	}
	ctx.JSON(http.StatusOK, file)
}

func (fs *FilesController) ListFilesByUser(ctx *gin.Context) {
	userId, authErr := authenticateUser(ctx)
	if authErr != nil {
		return
	}
	files, err := fs.service.ListFilesByUser(userId)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "file not found"})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get file"})
		return
	}
	ctx.JSON(http.StatusOK, files)
}

// update file metadata, not the content, that is managed by the files service
func (fs *FilesController) UpdateFile(ctx *gin.Context) {
	// TODO
}

func (fs *FilesController) DeleteFile(ctx *gin.Context) {
	// TODO
}
