package v1

import (
	"github.com/gin-gonic/gin"

	files "github.com/kamil-abbasi/CloudFileOperationsBackend/internal/files/controllers/v1"
)

func NewRouter(r *gin.Engine) {
	v1Router := r.Group("/v1")
	{
		filesRouter := v1Router.Group("/files")
		{
			controller := files.NewController()

			filesRouter.GET("/:id", controller.FindOne)

			filesRouter.GET("/:id/download", controller.Download)

			filesRouter.POST("", controller.Upload)

			filesRouter.PATCH("/:id", controller.Update)

			filesRouter.DELETE("/:id", controller.Remove)
		}
	}
}
