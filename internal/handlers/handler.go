package handlers

import (
	"github.com/gin-gonic/gin"
)

type Handlers interface {
	Register(router *gin.Engine,user interface{})
}
