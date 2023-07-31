package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"

	"go-monorepo-boilerplate/exceptions"
	"go-monorepo-boilerplate/helpers"
	"go-monorepo-boilerplate/middleware"
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

	err := c.Bind(&request)
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
		accessToken, refreshToken, err := middleware.GenerateJwt(result.Id, result.Name, result.Username, result.Email, "")
		if err != nil {
			logrus.Errorf("%w: %v", exceptions.InternalServerError, err)
			c.JSON(http.StatusInternalServerError, helpers.ErrorResponse{
				Message: err.Error(),
				Success: false,
				Error:   err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, helpers.SuccessResponse{
			Data: map[string]interface{}{
				"user": map[string]string{
					"id":       result.Id,
					"name":     result.Name,
					"username": result.Username,
					"email":    result.Email,
					"avatar":   "",
				},
				"accessToken":  accessToken,
				"refreshToken": refreshToken,
			},
			Message: "Successfully Login",
			Success: true,
		})
		return
	}
}
