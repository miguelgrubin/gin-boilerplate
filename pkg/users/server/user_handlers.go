package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/miguelgrubin/gin-boilerplate/pkg/users/usecases"
)

type UserHandlers struct {
	usecase usecases.UserUseCasesInterface
}

func NewUserHandlers(u usecases.UserUseCasesInterface) UserHandlers {
	return UserHandlers{
		usecase: u,
	}
}

func (uh *UserHandlers) UserCreateHandler(c *gin.Context) {
	var userParams UserCreateRequest
	if err := c.ShouldBindJSON(&userParams); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	user, err := uh.usecase.Creator(usecases.UserCreatorParams{
		Username:  userParams.Username,
		FirstName: userParams.FirstName,
		LastName:  userParams.LastName,
		Email:     userParams.Email,
		Password:  userParams.Password,
		Phone:     userParams.Phone,
	})

	if err != nil {
		handleError(c, err)
		return
	}
	c.JSON(http.StatusCreated, UserResponseFromDomain(user))
}

func (uh *UserHandlers) UserShowHandler(c *gin.Context) {
	username := c.Param("username")
	user, err := uh.usecase.Showher(username)

	if err != nil {
		handleError(c, err)
		return
	}
	c.JSON(http.StatusOK, UserResponseFromDomain(user))
}

func (uh *UserHandlers) UserUpdateHandler(c *gin.Context) {
	username := c.Param("username")
	var userParams UserUpdateRequest
	if err := c.ShouldBindJSON(&userParams); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	payload := (usecases.UserUpdatersParams)(userParams)
	user, err := uh.usecase.Updater(username, payload)

	if err != nil {
		handleError(c, err)
		return
	}
	c.JSON(http.StatusOK, UserResponseFromDomain(user))
}

func (uh *UserHandlers) UserDeleteHandler(c *gin.Context) {
	username := c.Param("username")
	err := uh.usecase.Deleter(username)

	if err != nil {
		handleError(c, err)
		return
	}
	c.JSON(http.StatusNoContent, nil)
}

func (uh *UserHandlers) UserLoginHandler(c *gin.Context) {
	var loginParams UserLoginRequest
	if err := c.ShouldBindJSON(&loginParams); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	token, refreshToken, err := uh.usecase.LoggerIn(loginParams.Username, loginParams.Password)

	if err != nil {
		handleError(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"token":         token,
		"refresh_token": refreshToken,
	})
}

func (uh *UserHandlers) UserRefreshHandler(c *gin.Context) {
	token, refreshToken, err := uh.usecase.RefreshToken(c.GetHeader("Authorization"))
	if err != nil {
		handleError(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"token":         token,
		"refresh_token": refreshToken,
	})
}

func (uh *UserHandlers) UserLogoutHandler(c *gin.Context) {
	err := uh.usecase.LoggerOut(c.GetHeader("Authorization"))
	if err != nil {
		handleError(c, err)
		return
	}
	c.JSON(http.StatusNoContent, nil)
}

func (uh *UserHandlers) SetupRoutes(r *gin.RouterGroup) {
	r.POST("/users", uh.UserCreateHandler)
	r.GET("/user/:username", uh.UserShowHandler)
	r.PATCH("/user/:username", uh.UserUpdateHandler)
	r.DELETE("/user/:username", uh.UserDeleteHandler)
	r.POST("/auth/login", uh.UserLoginHandler)
	r.POST("/auth/refresh", uh.UserRefreshHandler)
	r.POST("/auth/logout", uh.UserLogoutHandler)
}
