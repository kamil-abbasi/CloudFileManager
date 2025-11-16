package internal

import (
	"github.com/gin-gonic/gin"

	v1 "github.com/kamil-abbasi/CloudFileOperationsBackend/internal/api/v1"
)

func NewApiRouter() *gin.Engine {
	r := gin.Default()

	v1.NewRouter(r)

	return r
}
