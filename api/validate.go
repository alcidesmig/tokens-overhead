package api

import (
	"log"
	"time"
	"tokens-overhead/service"

	"github.com/gin-gonic/gin"
)

type validateHandler struct {
	svc service.TokenService
}

func NewValidateHandler(
	svc service.TokenService,
) *validateHandler {
	return &validateHandler{svc: svc}
}

func (v *validateHandler) Token(c *gin.Context) {
	log.Printf("serving... %s", time.Now())
	token := c.Request.Header["Authorization"]
	if len(token) != 1 {
		c.JSON(400, gin.H{
			"message": "bad request",
		})
		return
	}
	valid, err := v.svc.SleepAndValidate(token[0], 10)
	if err != nil {
		c.JSON(401, gin.H{
			"message": "error validating token",
		})
		return
	}
	c.JSON(200, gin.H{
		"valid": valid,
	})
}
