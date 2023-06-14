package routes

import (
	"database/sql"

	"github.com/TomyPY/FinTracker/cmd/api/handlers"
	"github.com/TomyPY/FinTracker/docs"
	"github.com/TomyPY/FinTracker/internal/user"
	"github.com/TomyPY/FinTracker/internal/wallet"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Router interface {
	MapRoutes()
}

type router struct {
	db  *sql.DB
	eng *gin.Engine
	rg  *gin.RouterGroup
}

func NewRouter(db *sql.DB, eng *gin.Engine) Router {
	return &router{
		db:  db,
		eng: eng,
		rg:  eng.Group("/api/v1"),
	}
}
func (r *router) MapRoutes() {

	r.buildHealthRoutes()
	r.buildWalletRoutes()
	r.buildUserRoutes()
	r.buildSwaggerRoutes()

}

func (r *router) buildHealthRoutes() {
	r.eng.GET("/ping", handlers.PingHandler())
}

func (r *router) buildWalletRoutes() {

	repository := wallet.NewRepository(r.db)
	service := wallet.NewService(repository)
	wHandler := handlers.NewWallet(service)

	wRouter := r.rg.Group("/wallets")

	{
		wRouter.GET("/", wHandler.GetAll())
		wRouter.GET("/:id", wHandler.Get())
		wRouter.GET("/:id/transactions", wHandler.GetReportTransactions())
		wRouter.POST("/:id/add", wHandler.AddMoney())
		wRouter.POST("/:id/substract", wHandler.SubstractMoney())
		wRouter.DELETE("/:id", wHandler.Delete())
	}
}

func (r *router) buildUserRoutes() {

	repository := user.NewRepository(r.db)
	service := user.NewService(repository)
	userHandler := handlers.NewUser(service)

	usersRouter := r.rg.Group("/users")
	{
		usersRouter.GET("/", userHandler.GetAll())
		usersRouter.GET("/:id", userHandler.Get())
		usersRouter.POST("/", userHandler.Create())
		usersRouter.PATCH("/:id", userHandler.Update())
		usersRouter.DELETE("/:id", userHandler.Delete())

	}

}

func (r *router) buildSwaggerRoutes() {
	docs.SwaggerInfo.Host = "localhost:8080"
	r.rg.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
