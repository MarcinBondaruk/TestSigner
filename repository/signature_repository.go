package repository

import "github.com/MarcinBondaruk/TestSigner/model"

type SignatureRepository interface {
	Add(st model.SignedTest) error
	Retrieve(userId, signature string) (model.SignedTest, error)
}
