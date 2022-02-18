package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/swanden/service-template/internal/domain/user/usecase"
	"github.com/swanden/service-template/internal/infrastructure/domain/user/repository"
	"github.com/swanden/service-template/pkg/logger"
	"net/http"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "github.com/swanden/service-template/docs"
)

// Swagger spec:
// @title       Service Template
// @description Service Template
// @version     1.0
// @host        localhost:8000
// @BasePath    /v1
func NewRouter(handler *gin.Engine, userRepository *repository.UserRepository, useCase *usecase.UserUseCase, validate *validator.Validate, log logger.Interface) {
	handler.Use(gin.Logger())
	handler.Use(gin.Recovery())

	//swaggerHandler := ginSwagger.DisablingWrapHandler(swaggerFiles.Handler, "DISABLE_SWAGGER_HTTP_HANDLER")
	handler.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	h := handler.Group("/v1")
	{
		h.GET("/", index)
		newUserController(h, userRepository, *useCase, validate, log)
	}
}

type apiInfo struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}

// @Summary     Show API info
// @Description Show API info
// @ID          index
// @Tags  	    api info
// @Accept      json
// @Produce     json
// @Success     200 {string} apiInfo
// @Failure     500
// @Router      / [get]
func index(c *gin.Context) {
	info := apiInfo{
		Name:    "API",
		Version: "1.0",
	}

	c.JSON(http.StatusOK, info)
}
