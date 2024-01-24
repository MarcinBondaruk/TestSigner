package main

import (
	"net/http"
	"time"

	"github.com/MarcinBondaruk/TestSigner/config"
	"github.com/MarcinBondaruk/TestSigner/controller"
	"github.com/MarcinBondaruk/TestSigner/model"
	"github.com/MarcinBondaruk/TestSigner/repository"
	"github.com/MarcinBondaruk/TestSigner/router"
	"github.com/MarcinBondaruk/TestSigner/service"
)

func main() {
	db := config.GetDbConnection()
	db.Table("signed_tests").AutoMigrate(&model.SignedTest{})

	signedTestsRepo := repository.NewPostgreSignatureRepository(db)
	signerService := service.NewSignerService(signedTestsRepo)
	signerController := controller.NewTestSignerController(signerService)
	router := router.NewRouter(signerController)

	// basiclly taken from https://pkg.go.dev/net/http
	server := &http.Server{
		Addr:           ":8080", // should be read by os.Getenv()
		Handler:        router,
		ReadTimeout:    5 * time.Second,
		WriteTimeout:   5 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
