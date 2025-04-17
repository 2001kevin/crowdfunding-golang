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
        
        errors := helper.FormatValidationError(err)
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

func (h *userHandler) Login(c *gin.Context){
    var input user.LoginInput

    err := c.ShouldBindJSON(&input)

    if  err != nil {
        errors := helper.FormatValidationError(err)
        errorMessage := gin.H{"errors":errors}

        response := helper.APIResponse("Login failed", http.StatusUnprocessableEntity, "Failed", errorMessage)
        c.JSON(http.StatusBadRequest, response)
        return
    }

    loggedinUser, err := h.userService.LoginInput(input)

    if err != nil {
        errorMessage := gin.H{"errors":err.Error()}
        response := helper.APIResponse("Login failed", http.StatusUnprocessableEntity, "Failed", errorMessage)
        c.JSON(http.StatusBadRequest, response)
        return
    }

    formatter := user.FormatUser(loggedinUser, "token acak")

    response := helper.APIResponse("Successfully Loggedin", http.StatusOK, "success", formatter)
    c.JSON(http.StatusOK, response)
}

func(h *userHandler) CheckEmailAvailability(c *gin.Context){
    var input user.CheckEmailInput

    err := c.ShouldBindJSON(&input)
    if  err != nil {
        errors := helper.FormatValidationError(err)
        errorMessage := gin.H{"errors":errors}

        response := helper.APIResponse("Checking email failed", http.StatusUnprocessableEntity, "Failed", errorMessage)
        c.JSON(http.StatusUnprocessableEntity, response)
        return
    }

    isEmailAvailable, err := h.userService.IsEmailAvailable(input)
    if err != nil{
        errorMessage := gin.H{"errors": "Server Error"}

        response := helper.APIResponse("Checking email failed", http.StatusUnprocessableEntity, "Failed", errorMessage)
        c.JSON(http.StatusUnprocessableEntity, response)
        return
    }

    data := gin.H{
        "is_available" : isEmailAvailable,
    }

    metaMessage := "Email has been registered"

    if isEmailAvailable {
        metaMessage = "Email is available"
    }
    
    response := helper.APIResponse(metaMessage, http.StatusOK, "success", data)
    c.JSON(http.StatusOK, response)
}
