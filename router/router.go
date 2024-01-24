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
	tests := apiv1.Group("/tests")
	// it is restful if opration is the last part of an url
	tests.POST("/sign", testSignerController.Sign)
	// basiclly retrieve all but should be 1, could be refactored easily
	tests.GET("", testSignerController.RetrieveByUserIdAndSignature)

	return testSignerService
}
