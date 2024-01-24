package service

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"strings"
	"time"

	"github.com/MarcinBondaruk/TestSigner/api/request"
	"github.com/MarcinBondaruk/TestSigner/model"
	"github.com/MarcinBondaruk/TestSigner/repository"
)

func calculateSignature(userId string, reqData []request.SignAnswersRequest) string {
	toBeHashed := userId

	for _, v := range reqData {
		toBeHashed += v.Question
	}

	data := []byte(toBeHashed)
	hash := md5.Sum(data)

	return hex.EncodeToString(hash[:])
}

// helper function
func base64Answers(data []request.SignAnswersRequest) string {
	var answers []string

	for _, v := range data {
		answers = append(answers, base64.StdEncoding.EncodeToString([]byte(v.Answer)))
	}

	return strings.Join(answers, ",")
}

// helper function
func base64Questions(data []request.SignAnswersRequest) string {
	var questions []string

	for _, v := range data {
		questions = append(questions, base64.StdEncoding.EncodeToString([]byte(v.Question)))
	}

	return strings.Join(questions, ",")
}

func deBase64Answers(based string) string {
	var answers []string
	var bytes []byte
	tmp := strings.Split(based, ",")

	for i := 0; i < len(tmp); i++ {
		bytes, _ = base64.StdEncoding.DecodeString(tmp[i])
		answers = append(answers, string(bytes))
	}

	return strings.Join(answers, ",")
}

type SignerServiceImpl struct {
	signatureRepository repository.SignatureRepository
}

func NewSignerService(signatureRepository repository.SignatureRepository) SignerService {
	return &SignerServiceImpl{
		signatureRepository: signatureRepository,
	}
}

func (s SignerServiceImpl) Sign(userId string, reqData []request.SignAnswersRequest) (string, error) {

	signature := calculateSignature(userId, reqData)
	questionsStringified := base64Questions(reqData)
	answersStringified := base64Answers(reqData)

	newSignedTest := model.SignedTest{
		UserId:    userId,
		Answers:   answersStringified,
		Questions: questionsStringified,
		Signature: calculateSignature(userId, reqData),
		Timestamp: time.Now(),
	}

	err := s.signatureRepository.Add(newSignedTest)

	if err != nil {
		return "", err
	}

	return signature, nil
}

// basiclly a proxy for signed tests repository
func (s SignerServiceImpl) Retrieve(userId, signature string) (string, time.Time, error) {
	signedTest, err := s.signatureRepository.Retrieve(userId, signature)

	if err != nil {
		return "", time.Time{}, err
	}

	return deBase64Answers(signedTest.Answers), signedTest.Timestamp, nil
}
