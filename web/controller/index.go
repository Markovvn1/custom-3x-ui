package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type IndexController struct {
	BaseController
}

func NewIndexController(g *gin.RouterGroup) *IndexController {
	a := &IndexController{}
	a.initRouter(g)
	return a
}

func (a *IndexController) initRouter(g *gin.RouterGroup) {
	g.GET("/", a.index)
}

func (a *IndexController) index(c *gin.Context) {
	c.Redirect(http.StatusTemporaryRedirect, "panel/inbounds")
}
