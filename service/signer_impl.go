package service

import (
	"time"

	"github.com/MarcinBondaruk/TestSigner/repository"
)

type SignerServiceImpl struct {
	signatureRepository repository.SignatureRepository
}

func NewSignerService(signatureRepository repository.SignatureRepository) SignerService {
	return &SignerServiceImpl{
		signatureRepository: signatureRepository,
	}
}

func (s SignerServiceImpl) Sign(data string) (string, error) {
	// validate request data
	// create new obj
	// calculate hash
	// add to repository
	return "", nil
}

// basiclly a proxy for signed tests repository
func (s SignerServiceImpl) Retrieve(userId, signature string) (bool, string, time.Time) {
	signedTest, err := s.signatureRepository.Retrieve(userId, signature)

	if err != nil {
		panic(err) // for now
	}

	return true, signedTest.Answers, signedTest.Timestamp
}
