package controller

import "time"

type RetrieveByUserIdAndSignatureResponse struct {
	Status    string
	Answers   string
	Timestamp time.Time
}
