package links

import (
	"github.com/gin-gonic/gin"
)

// links CRUD -- functions in controllers.go
func Routes(g *gin.RouterGroup) {
	g.GET("/",
		GetLinks)
	g.POST("/:key",
		PostNewLinkByKey)
	g.GET("/:key",
		GetLinkByKey)
	g.PUT("/:key",
		UpdateLinkByKey)
	g.DELETE("/:key",
		DeleteLinkByKey)
	g.PUT("/:key/active",
		ActiveToggleByKey)
	g.POST("/restore",
		PostDumpRestore)
}
