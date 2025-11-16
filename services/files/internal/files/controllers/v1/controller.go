package files

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type FilesController struct{}

func NewController() FilesController {
	return FilesController{}
}

// GET /v1/files/test
func (controller *FilesController) Test(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello, World!",
	})
}

// POST /v1/files
func (controller *FilesController) Upload(c *gin.Context) {}

// GET /v1/files/:id/download
func (controller *FilesController) Download(c *gin.Context) {}

// DELETE /v1/files/:id
func (controller *FilesController) Remove(c *gin.Context) {}

// GET /v1/files/:id
func (controller *FilesController) FindOne(c *gin.Context) {}

// PATCH /v1/files/:id
func (controller *FilesController) Update(c *gin.Context) {}
