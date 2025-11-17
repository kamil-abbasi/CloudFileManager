package files

import (
	"errors"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/kamil-abbasi/CloudFileOperationsBackend/internal/files"
	"github.com/kamil-abbasi/CloudFileOperationsBackend/internal/utils"
)

type FilesController struct{}

func NewController() FilesController {
	return FilesController{}
}

// POST /v1/files
func (controller *FilesController) Upload(c *gin.Context) {
	rawFile, err := c.FormFile("file")
	location := c.PostForm("location")

	if location == "" {
		location = "/"
	}

	if err != nil {
		if errors.Is(err, http.ErrMissingFile) {
			c.JSON(http.StatusBadRequest, &utils.HttpError{
				Code:    http.StatusBadRequest,
				Message: "File missing",
			})
		} else {
			c.JSON(http.StatusInternalServerError, &utils.HttpError{
				Code:    http.StatusInternalServerError,
				Message: "Unknown error",
			})
		}
	}

	id, err := uuid.NewUUID()

	if err != nil {
		c.JSON(http.StatusInternalServerError, &utils.HttpError{
			Code:    http.StatusInternalServerError,
			Message: "Error while generating file id",
		})
	}

	file := &files.File{
		Id:       id.String(),
		Filename: rawFile.Filename,
		Location: "/data" + location + "/",
		Size:     uint64(rawFile.Size),
	}

	log.Printf("%+v\n", file)

	c.SaveUploadedFile(rawFile, file.Location+file.Filename)

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
