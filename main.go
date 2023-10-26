package main

import (
	"fmt"
	"net/http"
	"time"

	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
	"github.com/nandes007/employee-claim-request/app"
	"github.com/nandes007/employee-claim-request/config"
	"github.com/nandes007/employee-claim-request/controller"
	"github.com/nandes007/employee-claim-request/helper"
	"github.com/nandes007/employee-claim-request/repository"
	"github.com/nandes007/employee-claim-request/service"
)

var (
	uni      *ut.UniversalTranslator
	validate *validator.Validate
)

func main() {
	config.Load()
	db := app.NewDB()

	validate := validator.New()

	userRepository := repository.NewUserRepository()
	userService := service.NewUserService(userRepository, db, validate)
	userController := controller.NewUserController(userService)

	companyRepository := repository.NewCompanyRepository()
	companyService := service.NewCompanyService(companyRepository, db, validate)
	companyController := controller.NewCompanyController(companyService)

	authRepository := repository.NewAuthRepository()
	authService := service.NewAuthService(authRepository, companyRepository, db, validate)
	authController := controller.NewAuthController(authService)

	claimRequestRepository := repository.NewClaimRequestRepository()
	claimRequestService := service.NewClaimRequestService(claimRequestRepository, userRepository, db, validate)
	claimRequestController := controller.NewClaimRequestController(claimRequestService)

	router := app.NewRouter(authController, companyController, claimRequestController, userController)

	server := http.Server{
		Addr:        ":8080",
		Handler:     router,
		ReadTimeout: 5 * time.Second,
	}

	fmt.Printf("Server is running at http://localhost%s\n", server.Addr)
	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
