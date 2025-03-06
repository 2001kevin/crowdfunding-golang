package handler

import (
	"crowdfunding/helper"
	"crowdfunding/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userService user.Service
}

func NewUserHandler(userService user.Service) *userHandler{
    return &userHandler{userService}
}

func (h *userHandler) RegisterUser(c *gin.Context){
    var input user.RegisterUserInput

    err := c.ShouldBindJSON(&input)
    if err != nil{
        
        errors := helper.FormatError(err)
        
        errorMessage := gin.H{"errors":errors}

        response := helper.APIResponse("Registered accound failed", http.StatusUnprocessableEntity, "Failed", errorMessage)
        c.JSON(http.StatusBadRequest, response)
        return
    }

    newUser, err := h.userService.RegisterUser(input)
    if err != nil{
        response := helper.APIResponse("Registered accound failed", http.StatusBadRequest, "Failde", nil)
        c.JSON(http.StatusBadRequest, response)
        return
    }

    formatter := user.FormatUser(newUser, "token token")

    response := helper.APIResponse("Account has been registered", http.StatusOK, "success", formatter)

    c.JSON(http.StatusOK, response)
}