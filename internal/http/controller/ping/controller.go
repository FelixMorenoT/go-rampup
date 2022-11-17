package ping

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Controller struct {
}

func NewController() *Controller {
	return &Controller{}
}

func (ctrl *Controller) Ping(c *gin.Context) {
	c.JSON(http.StatusOK, "Pong")
}
