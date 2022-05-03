package api

import (
	"fmt"
	"tokens-overhead/service"

	"github.com/gin-gonic/gin"
)

func InitAPI(svc service.TokenService, apiHost, apiPort string) {
	r := gin.Default()
	r.GET("/healthcheck", HealthCheck)

	handler := NewValidateHandler(svc)

	r.GET("/token", handler.Token)

	r.Run(fmt.Sprintf("%s:%s", apiHost, apiPort))
}
