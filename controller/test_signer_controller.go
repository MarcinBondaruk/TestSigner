package controller

import (
	"net/http"

	"github.com/MarcinBondaruk/TestSigner/api/response"
	"github.com/MarcinBondaruk/TestSigner/service"
	"github.com/gin-gonic/gin"
)

type TestSignerController struct {
	signerService service.SignerService
}

func NewTestSignerController(signerService service.SignerService) *TestSignerController {
	return &TestSignerController{}
}

func (tsc TestSignerController) Sign(ctx *gin.Context) {
	signature, err := tsc.signerService.Sign("SOME DATA")

	if err != nil {
		panic(err)
	}

	ctx.JSON(http.StatusOK, response.SignResponse{Signature: signature})
}

func (tsc TestSignerController) RetrieveByUserIdAndSignature(ctx *gin.Context) {
	succ, answers, timestamp := tsc.signerService.Retrieve("someid", "somesignature")

	if succ {
		ctx.JSON(200, response.RetrieveByUserIdAndSignatureResponse{
			Status:    "OK",
			Answers:   answers,
			Timestamp: timestamp,
		})
	} else {
		ctx.JSON(500, "some error")
	}
}
