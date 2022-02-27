package api

import (
	"main/db"

	"github.com/gin-gonic/gin"
)

// Setup - call this method to setup routes
func Setup(router *gin.Engine) {

	db.SetupDB()
	// CORSMiddleware()
	SetupLottoAPI(router)
	SetupAuthenAPI(router)
	SetupProductAPI(router)
	SetupTransactionAPI(router)

}
