package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ryosantouchh/7solutions/piefiredire/internal/core/ports"
)

type GinContext struct {
	// embedded field for gin context
	*gin.Context
}

type Context struct {
	http.ResponseWriter
	*http.Request
}

func NewApiContext(ctx *gin.Context) *GinContext {
	return &GinContext{Context: ctx}
}

func NewHTTPContext(w http.ResponseWriter, r *http.Request) *Context {
	return &Context{}
}

func (c *GinContext) JSON(statusCode int, response interface{}) {
	c.Context.JSON(statusCode, response)
}

func (c *Context) JSON(statusCode int, rawData interface{}) {
	jsonData, err := json.Marshal(rawData)
	if err != nil {
		_ = fmt.Errorf("Error encoding data to JSON: %v \n", err)
		return
	}

	c.ResponseWriter.Header().Set("Content-Type", "application/json")
	_, err = c.ResponseWriter.Write(jsonData)
	if err != nil {
		_ = fmt.Errorf("Error writing json response: %v \n", err)
		return
	}
}

func GinHandler(handler func(ports.HTTPContext)) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		handler(NewApiContext(ctx))
	}
}

func HttpHandler(handler func(ports.HTTPContext)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		handler(NewHTTPContext(w, r))
	}
}
