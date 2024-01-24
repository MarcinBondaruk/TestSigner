package controller

import (
	"net/http"

	"github.com/MarcinBondaruk/TestSigner/api/request"
	"github.com/MarcinBondaruk/TestSigner/api/response"
	"github.com/MarcinBondaruk/TestSigner/service"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type TestSignerController struct {
	signerService service.SignerService
}

func isUUID(s string) bool {
	_, err := uuid.Parse(s)
	return err == nil
}

func NewTestSignerController(signerService service.SignerService) *TestSignerController {
	return &TestSignerController{signerService: signerService}
}

func (tsc *TestSignerController) Sign(ctx *gin.Context) {
	var signRequest []request.SignAnswersRequest
	userId, _ := ctx.Get("userId")
	// i will just assume its always string while setting
	usrIdVal, _ := userId.(string)

	err := ctx.ShouldBindJSON(&signRequest)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	signature, err := tsc.signerService.Sign(usrIdVal, signRequest)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, response.SignResponse{Signature: signature})
}

func (tsc *TestSignerController) RetrieveByUserIdAndSignature(ctx *gin.Context) {
	userId := ctx.Query("user_id")
	signature := ctx.Query("signature")

	if userId == "" || signature == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "URL params: user_id and signature are required"})
		return
	}

	if !isUUID(userId) {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "user_id must be valid UUID"})
		return
	}

	answers, timestamp, err := tsc.signerService.Retrieve(userId, signature)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "OK", "answers": answers, "timestamp": timestamp})
}
