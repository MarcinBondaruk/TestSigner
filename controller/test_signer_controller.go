package controller

import (
	"net/http"

	"github.com/MarcinBondaruk/TestSigner/api/request"
	"github.com/MarcinBondaruk/TestSigner/api/response"
	"github.com/MarcinBondaruk/TestSigner/service"
	"github.com/gin-gonic/gin"
)

type TestSignerController struct {
	signerService service.SignerService
}

func NewTestSignerController(signerService service.SignerService) *TestSignerController {
	return &TestSignerController{signerService: signerService}
}

func (tsc *TestSignerController) Sign(ctx *gin.Context) {
	var signRequest []request.SignAnswersRequest
	userId := "018d3c5a-6d2b-79ed-94ab-90e9d9ba526b"
	err := ctx.ShouldBindJSON(&signRequest)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
	}

	signature, err := tsc.signerService.Sign(userId, signRequest)

	// handle signer error
	if err != nil {
		panic(err)
	}

	ctx.JSON(http.StatusOK, response.SignResponse{Signature: signature})
}

func (tsc *TestSignerController) RetrieveByUserIdAndSignature(ctx *gin.Context) {
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
