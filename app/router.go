package app

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/nandes007/employee-claim-request/controller"
	"github.com/nandes007/employee-claim-request/exception"
	"github.com/nandes007/employee-claim-request/middleware"
)

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Welcome!\n")
}

func NewRouter(authController controller.AuthController, companyController controller.CompanyController, claimRequestController controller.ClaimRequestController, userController controller.UserController) *httprouter.Router {
	router := httprouter.New()
	router.GET("/", Index)
	router.POST("/api/admin/register", authController.AdminRegister)
	router.POST("/api/register", authController.Register)
	router.POST("/api/login", authController.Login)
	router.PUT("/api/users", middleware.JwtAuthMiddleware(userController.Update))

	// Company
	router.GET("/api/companies", companyController.GetAll)
	router.GET("/api/companies/:id", companyController.Find)
	router.POST("/api/companies", companyController.Create)
	router.PUT("/api/companies/:id", companyController.Update)
	router.DELETE("/api/companies/:id", companyController.Delete)

	// Claim Request
	router.GET("/api/claim-requests", middleware.JwtAuthMiddleware(claimRequestController.GetAll))
	router.GET("/api/claim-requests/:id", middleware.JwtAuthMiddleware(claimRequestController.Find))
	router.POST("/api/claim-requests", middleware.JwtAuthMiddleware(claimRequestController.Create))
	router.PUT("/api/claim-requests/:id", middleware.JwtAuthMiddleware(claimRequestController.Update))
	router.PUT("/api/claim-requests/:id/status", middleware.JwtAuthMiddleware(claimRequestController.UpdateStatus))
	router.DELETE("/api/claim-requests/:id", middleware.JwtAuthMiddleware(claimRequestController.Delete))
	router.POST("/api/upload", middleware.JwtAuthMiddleware(claimRequestController.UploadFile))

	router.PanicHandler = exception.ErrorHandler

	return router
}
