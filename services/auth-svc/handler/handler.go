package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"go-monorepo-boilerplate/helpers"
	"go-monorepo-boilerplate/services/auth-svc/app/usecase/user"
	"go-monorepo-boilerplate/services/auth-svc/entity"
)

type RestHandler struct {
	userUseCase user.IUseCase
}

func New(userUseCase user.IUseCase) IRestHandler {
	return &RestHandler{userUseCase}
}

// Login implements IRestHandler.
func (h *RestHandler) Login(c *gin.Context) {
	var request entity.AuthRequest

	err := c.ShouldBindJSON(&request)
	if err != nil {
		if len(request.Password) < 8 {
			c.JSON(http.StatusInternalServerError, helpers.ErrorResponse{
				Message: "Password min 8 character",
				Success: false,
			})
			return
		} else {
			c.JSON(http.StatusInternalServerError, helpers.ErrorResponse{
				Message: "Email or Password Not Match",
				Success: false,
			})
			return
		}
	}

	result, err := h.userUseCase.Login(request.Email, request.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, helpers.ErrorResponse{
			Message: err.Error(),
			Success: false,
			Error:   err.Error(),
		})
		return
	} else {
		c.JSON(http.StatusOK, helpers.SuccessResponse{
			Data:    result,
			Message: "Successfully Login",
			Success: true,
		})
		return
	}
}
