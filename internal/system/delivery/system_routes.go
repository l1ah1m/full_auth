package systemDelivery

import (
	"github.com/Point-AI/backend/config"
	"github.com/Point-AI/backend/internal/system/delivery/controller"
	"github.com/Point-AI/backend/internal/system/infrastructure/client"
	"github.com/Point-AI/backend/middleware"
	"github.com/labstack/echo/v4"
	"github.com/minio/minio-go/v7"
	"go.mongodb.org/mongo-driver/mongo"
)

func RegisterSystemRoutes(e *echo.Echo, cfg *config.Config, db *mongo.Database, str *minio.Client) {
	systemGroup := e.Group("/system")

	strc := client.NewStorageClientImpl(str)

	sc := controller.NewSystemController(, cfg, str)

	projectGroup := systemGroup.Group("/project")
	projectGroup.POST("", sc.CreateProject, middleware.ValidateAccessTokenMiddleware(cfg.Auth.JWTSecretKey))
	projectGroup.POST("/:member", sc.AddProjectMember, middleware.ValidateAccessTokenMiddleware(cfg.Auth.JWTSecretKey))
	projectGroup.GET("/:id", sc.GetProjectByID, middleware.ValidateAccessTokenMiddleware(cfg.Auth.JWTSecretKey))
	projectGroup.GET("", sc.GetAllProjects, middleware.ValidateAccessTokenMiddleware(cfg.Auth.JWTSecretKey))
	projectGroup.PUT("/:member", sc.UpdateProjectMember, middleware.ValidateAccessTokenMiddleware(cfg.Auth.JWTSecretKey))
	projectGroup.PUT("/:id", sc.UpdateProjectByID, middleware.ValidateAccessTokenMiddleware(cfg.Auth.JWTSecretKey))
	projectGroup.PUT("/leave/:id", sc.LeaveProject, middleware.ValidateAccessTokenMiddleware(cfg.Auth.JWTSecretKey))
	projectGroup.DELETE("/:member", sc.DeleteProjectMember, middleware.ValidateAccessTokenMiddleware(cfg.Auth.JWTSecretKey))
	projectGroup.DELETE("/:id", sc.DeleteProjectByID, middleware.ValidateAccessTokenMiddleware(cfg.Auth.JWTSecretKey))
}

