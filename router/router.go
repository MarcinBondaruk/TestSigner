package router

import (
	"net/http"

	"github.com/MarcinBondaruk/TestSigner/controller"
	"github.com/gin-gonic/gin"
)

func NewRouter(testSignerController *controller.TestSignerController) *gin.Engine {
	// setup service at root level
	testSignerService := gin.Default()
	testSignerService.GET("", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "Test Signer by Marcin Bondaruk")
	})

	// do i need it?
	// testSignerService.NoRoute(func(ctx *gin.Context) {
	// 	ctx.JSON(http.StatusNotFound, gin.H{})
	// })

	// setup service api
	apiv1 := testSignerService.Group("/api/v1")
	apiv1.POST("/sign", testSignerController.Sign)
	apiv1.GET("/retrieve", testSignerController.RetrieveByUserIdAndSignature)

	return testSignerService
}
