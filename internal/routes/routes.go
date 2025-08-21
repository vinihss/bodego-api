package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/vinihss/bodego-api/internal/infrastructure/database/repositories"
	http_interfaces_authentication "github.com/vinihss/bodego-api/internal/interfaces/http/authentcation"
	customeruse "github.com/vinihss/bodego-api/internal/usecases/customer"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	_ "github.com/vinihss/bodego-api/docs"
	"github.com/vinihss/bodego-api/internal/interfaces/http/customer"
	"github.com/vinihss/bodego-api/middlewares"
)

// SetupRoutes @title Aiqfome API
func SetupRoutes(router *gin.Engine) {

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	authController := http_interfaces_authentication.NewAuthenticationController()
	router.POST("/authenticate", http_interfaces_authentication.NewAuthenticationHandler(authController).Authenticate)
	db, _ := gorm.Open(postgres.Open("host=dpg-d2jbkfbuibrs73defhq0-a user=posgres password=6REY3Hz6IzIODENTegTSiaac7zkAOL5u dbname=bodego port=5432 sslmode=disable"), &gorm.Config{})

	authorized := router.Group("/")
	authorized.Use(middlewares.JWTAuth())
	{

		custRepo := repositories.NewCustomerRepository(db)
		createCustomerUC := customeruse.NewCreateCustomerUseCase(custRepo)
		deleteCustomerUC := customeruse.NewDeleteCustomerUseCase(custRepo)
		findCustomerUC := customeruse.NewFindCustomerUseCase(custRepo)
		updateCustomerUC := customeruse.NewUpdateCustomerUseCase(custRepo)

		custController := http_interfaces_customer.NewCustomerController(createCustomerUC, deleteCustomerUC, findCustomerUC, updateCustomerUC)
		custHandler := http_interfaces_customer.NewCustomerHandler(custController)
		RegisterCustomerRoutes(router, custHandler)

	}

}
