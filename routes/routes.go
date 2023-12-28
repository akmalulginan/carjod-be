package routes

import (
	"fmt"
	"net/http"
	"os"

	"github.com/akmalulginan/carjod-be/middleware"
	authH "github.com/akmalulginan/carjod-be/services/auth/handler"
	authU "github.com/akmalulginan/carjod-be/services/auth/usecase"
	matchH "github.com/akmalulginan/carjod-be/services/match/handler"
	matchR "github.com/akmalulginan/carjod-be/services/match/repository"
	matchU "github.com/akmalulginan/carjod-be/services/match/usecase"
	premiumH "github.com/akmalulginan/carjod-be/services/premium/handler"
	premiumU "github.com/akmalulginan/carjod-be/services/premium/usecase"
	txC "github.com/akmalulginan/carjod-be/services/tx/repository"
	uploadH "github.com/akmalulginan/carjod-be/services/upload/handler"
	userH "github.com/akmalulginan/carjod-be/services/user/handler"
	userR "github.com/akmalulginan/carjod-be/services/user/repository"
	userU "github.com/akmalulginan/carjod-be/services/user/usecase"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Setup(r *gin.Engine, db *gorm.DB) {
	r.Static("/public", "./public")
	r.Static("/file", "./file")

	r.Use(middleware.CORS())

	txCoordinator := txC.NewTxCoordinator(db)

	userRepository := userR.NewUserRepository(db)
	authUsecase := authU.NewAuthUsecase(userRepository)
	authH.NewAuthHandler(r.Group(""), authUsecase)

	v1 := r.Group("v1")
	v1.Use(middleware.Authorize(middleware.NewJWTAuthService()))

	premiumUsecase := premiumU.NewPremiumUsecase(userRepository)
	premiumH.NewPremiumHandler(r.Group(""), v1, premiumUsecase)

	userUsecase := userU.NewUserUsecase(userRepository)
	userH.NewUserHandler(v1, userUsecase)

	matchRepository := matchR.NewMatchRepository(db)
	matchUsecase := matchU.NewMatchUsecase(matchRepository, userRepository, txCoordinator)
	matchH.NewMatchHandler(v1, matchUsecase)

	uploadH.NewUploadHandler(v1)

	server := http.Server{
		Addr:    fmt.Sprintf(":%s", os.Getenv("APP_PORT")),
		Handler: r,
	}

	server.ListenAndServe()
}
