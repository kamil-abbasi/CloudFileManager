package files

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/kamil-abbasi/CloudFileOperationsBackend/internal/config"
	"github.com/kamil-abbasi/CloudFileOperationsBackend/internal/files"
	"github.com/kamil-abbasi/CloudFileOperationsBackend/internal/shared"
)

type FilesController struct {
	config  *config.Config
	service *files.FileService
}

func NewController(config *config.Config) FilesController {
	return FilesController{
		config:  config,
		service: files.NewService(),
	}
}

// POST /v1/files
func (controller *FilesController) Upload(c *gin.Context) {
	rawFile, err := c.FormFile("file")
	location := c.PostForm("location")

	userRoot := controller.config.RootPath + "/user-dev"

	if err != nil {
		c.JSON(http.StatusBadRequest, &shared.HttpError{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
	}

	id, err := uuid.NewUUID()

	if err != nil {
		c.JSON(http.StatusInternalServerError, &shared.HttpError{
			Code:    http.StatusInternalServerError,
			Message: "Internal server error",
		})
	}

	file := &files.File{
		Id:       id.String(),
		Filename: rawFile.Filename,
		Location: location + "/",
		Size:     uint64(rawFile.Size),
	}

	controller.service.Create(files.FileCreateDto{
		Id:       file.Id,
		Filename: file.Filename,
		Location: file.Location,
		Size:     file.Size,
	})

	c.SaveUploadedFile(rawFile, userRoot+file.Location+file.Filename)
	c.JSON(http.StatusCreated, file)
}

// GET /v1/files/:id/download
func (controller *FilesController) Download(c *gin.Context) {}

// DELETE /v1/files/:id
func (controller *FilesController) Remove(c *gin.Context) {}

// GET /v1/files/:id
func (controller *FilesController) FindOne(c *gin.Context) {}

// PATCH /v1/files/:id
func (controller *FilesController) Update(c *gin.Context) {}
