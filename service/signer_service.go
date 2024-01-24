package service

import (
	"time"

	"github.com/MarcinBondaruk/TestSigner/api/request"
)

type SignerService interface {
	Sign(userId string, reqData []request.SignAnswersRequest) (string, error)
	Retrieve(userid, signature string) (string, time.Time, error)
}
