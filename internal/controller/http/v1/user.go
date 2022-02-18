package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/pkg/errors"
	"github.com/swanden/service-template/internal/domain"
	"github.com/swanden/service-template/internal/domain/user/repository"
	"github.com/swanden/service-template/internal/domain/user/usecase"
	"github.com/swanden/service-template/pkg/logger"
	"net/http"
)

type userController struct {
	userRepository repository.UserRepositoryInterface
	userUseCase    usecase.UserUseCase
	validate       *validator.Validate
	logger         logger.Interface
}

func newUserController(ginHandler *gin.RouterGroup, userRepository repository.UserRepositoryInterface, userUseCase usecase.UserUseCase, validate *validator.Validate, log logger.Interface) {
	u := &userController{
		userRepository: userRepository,
		userUseCase:    userUseCase,
		validate:       validate,
		logger:         log,
	}

	h := ginHandler.Group("/user")
	{
		h.GET("/", u.all)
		h.POST("/", u.create)
		h.DELETE("/:id", u.delete)
	}
}

type allResponseUser struct {
	ID        string `json:"id"`
	Email     string `json:"email"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

type allResponse struct {
	Users []allResponseUser `json:"users"`
}

// @Summary     Show all users
// @Description Show all users
// @ID          all
// @Tags  	    users
// @Accept      json
// @Produce     json
// @Success     201 {object} allResponse
// @Failure     500
// @Router      /user [get]
func (u *userController) all(c *gin.Context) {
	users, err := u.userRepository.All()
	if err != nil {
		u.logger.Error(err, logger.NewField("destination", "http - v1 - userController - all"))
		errorResponse(c, http.StatusInternalServerError, "internal server error")
		return
	}

	var response allResponse
	for _, user := range users {
		response.Users = append(response.Users, allResponseUser{
			ID:        user.ID.Value,
			Email:     user.Email.Value,
			FirstName: user.Name.FirstName,
			LastName:  user.Name.LastName,
		})
	}

	c.JSON(http.StatusOK, response)
}

type showResponse struct {
	Message string `json:"message"`
}

func (u *userController) show(c *gin.Context) {
	fmt.Println("show")
	c.JSON(http.StatusOK, &showResponse{
		Message: "UserController::show()",
	})
}

type createRequest struct {
	Email     string `json:"email" binding:"required" validate:"required,email" example:"user@example.com"`
	Password  string `json:"password" binding:"required" validate:"required,min=3" example:"password"`
	FirstName string `json:"first_name" binding:"required" validate:"required" example:"John"`
	LastName  string `json:"last_name" binding:"required" validate:"required" example:"Doe"`
}

type createResponse struct {
	ID string `json:"id"`
}

// @Summary     Create user
// @Description Create user
// @ID          create
// @Tags  	    users
// @Accept      json
// @Produce     json
// @Param		user body createRequest true "User info"
// @Success     201 {object} createResponse
// @Failure     400
// @Failure     500
// @Router      /user [post]
func (u *userController) create(c *gin.Context) {
	var request createRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		u.logger.Error(err, logger.NewField("destination", "http - v1 - userController - create"))
		errorResponse(c, http.StatusBadRequest, "invalid request body")
		return
	}

	if err := u.validate.Struct(request); err != nil {
		e := err.(validator.ValidationErrors)
		u.logger.Error(e.Error(), logger.NewField("destination", "http - v1 - userController - create"))
		errorsResponse(c, http.StatusBadRequest, e)
		return
	}

	createDTO := usecase.CreateDTO{
		Email:     request.Email,
		Password:  request.Password,
		FirstName: request.FirstName,
		LastName:  request.LastName,
	}

	user, err := u.userUseCase.Create(createDTO)
	if errors.Cause(err) == domain.Error {
		errorResponse(c, http.StatusUnprocessableEntity, err.Error())
		return
	}
	if err != nil {
		errorResponse(c, http.StatusInternalServerError, "Internal server error")
		return
	}

	response := createResponse{
		ID: user.ID.Value,
	}

	c.JSON(http.StatusCreated, response)
}

// @Summary     Remove user
// @Description Remove user
// @ID          delete
// @Tags  	    users
// @Accept      json
// @Produce     json
// @Success     200
// @Failure     400
// @Failure     500
// @Router      /user/{id} [delete]
func (u *userController) delete(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		u.logger.Error("id param is required", logger.NewField("destination", "http - v1 - userController - delete"))
		errorResponse(c, http.StatusBadRequest, "id param is required")
		return
	}

	deleteDTO := usecase.DeleteDTO{ID: id}
	if err := u.userUseCase.Delete(deleteDTO); err != nil {
		u.logger.Error(err, logger.NewField("destination", "http - v1 - userController - delete"))
		errorResponse(c, http.StatusInternalServerError, "internal server error")
		return
	}

	c.AbortWithStatus(http.StatusOK)
}
