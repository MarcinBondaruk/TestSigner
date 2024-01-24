package service

import "time"

type SignerService interface {
	Sign(data string) (string, error)
	Retrieve(userid, signature string) (bool, string, time.Time)
}
