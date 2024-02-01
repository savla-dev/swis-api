package finance

import (
	"github.com/gin-gonic/gin"
)

// finance CRUD -- functions in controllers.go
func Routes(g *gin.RouterGroup) {
	g.GET("/",
		GetRootData)
	g.POST("/restore",
		PostDumpRestore)

	g.GET("/accounts/",
		GetAccounts)
	g.GET("/accounts/owner/:key",
		GetAccountByOwnerKey)
	g.POST("/accounts/:key",
		PostNewAccountByKey)
	//g.GET("/accounts/:key",
	//	GetAccountByKey)
	g.PUT("/accounts/:key",
		UpdateAccountByKey)
	g.DELETE("/accounts/:key",
		DeleteAccountByKey)

	g.GET("/taxes/:owner/:year",
		DoTaxesByOwner)
}
