package main

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB
var err error

type TestAnswer struct {
	Id        uint `gorm:"primaryKey"`
	UserId    string
	Questions string
	Answers   string
	Signature string
	Timestamp time.Time
}

func signAnswers(c *gin.Context) {
	// todo: extract from JWT
	userId := "some-id-form-jwt-token"

	var requestBody struct {
		Questions string
		Answers   string
	}

	err := c.BindJSON(&requestBody)

	if err != nil {
		c.JSON(400, gin.H{"error": "Bad request"})
		return
	}

	// todo: generate based on request data
	signature := "SuperAwesomeSignature"
	testAnswer := TestAnswer{
		UserId:    userId,
		Questions: requestBody.Questions,
		Answers:   requestBody.Answers,
		Signature: signature,
		Timestamp: time.Now(),
	}

	db.Create(&testAnswer)

	c.JSON(200, gin.H{"signature": signature})

}

func verifySignature(c *gin.Context) {
	var requestBody struct {
		UserId    string
		Signature string
	}

	if err := c.BindJSON(&requestBody); err != nil {
		c.JSON(400, gin.H{"error": "Bad Request"})
		return
	}

	var testAnswer TestAnswer
	result := db.Where("user_id = ? AND signature = ?", requestBody.UserId, requestBody.Signature).First(&testAnswer)

	if result.Error != nil {
		c.JSON(404, gin.H{"error": "Signature not found"})
		return
	}

	c.JSON(200, gin.H{
		"status":    "ok",
		"answers":   testAnswer.Answers,
		"timestamp": testAnswer.Timestamp,
	})
}

func main() {
	// dsn := "host=" + os.Getenv("DB_HOST") +
	// 	" port=" + os.Getenv("DB_PORT") +
	// 	" user=" + os.Getenv("DB_USER") +
	// 	" password=" + os.Getenv("DB_PASSWORD") +
	// 	" dbname=" + os.Getenv("DB_NAME")

	dsn := "host=localhost port=5432 user=postgres password=postgres dbname=test_signer"

	db, err = gorm.Open(postgres.Open(dsn))
	if err != nil {
		log.Fatal("failed to connect to db err:", err)
	}

	db.AutoMigrate(&TestAnswer{})

	router := gin.Default()
	router.POST("/sign", signAnswers)
	router.POST("/verify", verifySignature)
	router.Run("localhost:8080")
}
