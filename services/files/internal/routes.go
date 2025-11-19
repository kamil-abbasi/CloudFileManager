package internal

import (
	"github.com/gin-gonic/gin"

	v1 "github.com/kamil-abbasi/CloudFileOperationsBackend/internal/api/v1"
	"github.com/kamil-abbasi/CloudFileOperationsBackend/internal/config"
)

func NewApiRouter(config *config.Config) *gin.Engine {
	r := gin.Default()

	r.GET("/healthcheck")

	v1.NewRouter(r, config)

	return r
}
