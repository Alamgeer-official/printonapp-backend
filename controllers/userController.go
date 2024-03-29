package controllers

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"githuh.com/printonapp/models"
	"githuh.com/printonapp/repository"
	"githuh.com/printonapp/services"
	"githuh.com/printonapp/utils"
)

var userService = services.NewUserService(repository.NewUserRepo())

func Signup(ctx *gin.Context) {

	// bind signup data in user model
	var user models.User
	err := ctx.ShouldBind(&user)
	if err != nil {
		utils.ReturnError(ctx, err, http.StatusInternalServerError)
		return
	}

	//create user
	data, err := userService.CreateUser(user)
	if err != nil {
		utils.ReturnError(ctx, err, http.StatusInternalServerError)
		return
	}

	utils.ReturnResponse(ctx, data, http.StatusCreated)
}
func Login(ctx *gin.Context) {
	var user map[string]string
	err := ctx.ShouldBind(&user)
	if err != nil {
		utils.ReturnError(ctx, err, http.StatusBadRequest)
		return
	}
	data, err := userService.Login(user)
	if err != nil {
		utils.ReturnError(ctx, err, http.StatusBadRequest)
		return
	}
	utils.ReturnResponse(ctx, data, http.StatusOK)
}

func GetUsers(ctx *gin.Context) {
	data, err := userService.GetUsers(ctx)
	if err != nil {
		utils.ReturnError(ctx, err, http.StatusBadRequest)
		return
	}
	utils.ReturnResponse(ctx, data, http.StatusOK)

}
func GetUserById(ctx *gin.Context) {
	data, err := userService.GetUsers(ctx)
	if err != nil {
		utils.ReturnError(ctx, err, http.StatusBadRequest)
		return
	}
	utils.ReturnResponse(ctx, data, http.StatusOK)

}

func Test(ctx *gin.Context) {

	utils.ReturnResponse(ctx, "server is running", http.StatusOK)

}

func IsEmailExists(ctx *gin.Context) {
	email := ctx.Query("email")
	if email == "" {
		utils.ReturnError(ctx, errors.New("Email parameter is required"), http.StatusBadRequest)

		return
	}

	exists, err := userService.IsEmailExists(email)
	if err != nil {
		utils.ReturnError(ctx, errors.New("Failed to check email existence"), http.StatusInternalServerError)

		return
	}

	utils.ReturnResponse(ctx, exists, http.StatusOK)

}
