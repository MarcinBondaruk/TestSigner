package repository

import (
	"errors"
	"fmt"

	"github.com/MarcinBondaruk/TestSigner/model"
	"gorm.io/gorm"
)

type PostgreRepositoryImpl struct {
	dbConn *gorm.DB
}

func NewPostgreSignatureRepository(db *gorm.DB) SignatureRepository {
	return &PostgreRepositoryImpl{dbConn: db}
}

func (r PostgreRepositoryImpl) Add(st model.SignedTest) error {
	result := r.dbConn.Create(&st)

	if result.Error != nil {
		return fmt.Errorf("error creating new record %w", result.Error)
	}

	return nil
}

func (r PostgreRepositoryImpl) Retrieve(userId, signature string) (model.SignedTest, error) {
	var signedTest model.SignedTest

	result := r.dbConn.Where("user_id = ? AND signature = ?", userId, signature).First(&signedTest)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		// errors could be remapped to application layer errors
		return model.SignedTest{}, fmt.Errorf("can not find record by given criteria. userid: %s, signature: %s", userId, signature)
	}

	return signedTest, nil
}
