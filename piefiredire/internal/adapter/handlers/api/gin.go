package api

import (
	"github.com/gin-gonic/gin"
	"github.com/ryosantouchh/7solutions/piefiredire/internal/core/ports"
)

type GinContext struct {
	// embedded field for gin context
	*gin.Context
}

func NewApiContext(ctx *gin.Context) *GinContext {
	return &GinContext{Context: ctx}
}

func (c *GinContext) JSON(statusCode int, response interface{}) {
	c.Context.JSON(statusCode, response)
}

func GinHandler(handler func(ports.HTTPContext)) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		handler(NewApiContext(ctx))
	}
}
