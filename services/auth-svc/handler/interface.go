package handler

import "github.com/gin-gonic/gin"

type IRestHandler interface {
	Login(c *gin.Context)
}
