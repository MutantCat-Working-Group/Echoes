package router

import "github.com/gin-gonic/gin"

type RouterTemplate interface {
	PrepareRouter() error
	InitRouter(context *gin.Engine) error
	DestroyRouter() error
}
