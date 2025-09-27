package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/vinihss/bodego-api/config"
	_ "github.com/vinihss/bodego-api/docs"
	"github.com/vinihss/bodego-api/internal/domain/customer"
	"github.com/vinihss/bodego-api/internal/domain/tab"

	"github.com/vinihss/bodego-api/internal/infrastructure/database/repositories"
	http_interfaces_authentication "github.com/vinihss/bodego-api/internal/interfaces/http/authentcation"
	"github.com/vinihss/bodego-api/internal/interfaces/http/customer"
	customeruse "github.com/vinihss/bodego-api/internal/usecases/customer"
	"github.com/vinihss/bodego-api/middlewares"
)

func SetupRoutes(router *gin.Engine) {

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	authController := http_interfaces_authentication.NewAuthenticationController()
	router.POST("/authenticate", http_interfaces_authentication.NewAuthenticationHandler(authController).Authenticate)

	authorized := router.Group("/")
	authorized.Use(middlewares.JWTAuth())
	{

		var customerRepository customer.Repository = repositories.NewCustomerRepository(config.DB)
		createCustomerUC := customeruse.NewCreateCustomer(customerRepository)
		deleteCustomerUC := customeruse.NewDeleteCustomer(customerRepository)
		findCustomerUC := customeruse.NewFindCustomer(customerRepository)
		updateCustomerUC := customeruse.NewUpdateCustomer(customerRepository)

		custController := http_interfaces_customer.NewCustomerController(createCustomerUC, deleteCustomerUC, findCustomerUC, updateCustomerUC)
		custHandler := http_interfaces_customer.NewCustomerHandler(custController)
		RegisterCustomerRoutes(router, custHandler)
		var tabRepository tab.Repository = repositories.NewTabRepository(config.DB)
		var createTab = tab.NewCreateTab(tabRepository)
		// User CRUD (apenas sysadmin)
		/*var userRepository user.Repository = repositories.NewUserRepository(config.DB)
		var userCreateUC = useruse.NewCreateUser(userRepository)
		userHandler := http_interfaces_user.NewHandler(userCreateUC)
		userGroup := authorized.Group("")
		userGroup.Use(middlewares.OnlySysAdmin())
		http_interfaces_user.RegisterUserRoutes(userGroup, userHandler)*/
	}

}
