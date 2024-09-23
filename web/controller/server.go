package controller

import (
	"x-ui/web/service"

	"github.com/gin-gonic/gin"
)

type ServerController struct {
	BaseController

	serverService service.ServerService
}

func NewServerController(g *gin.RouterGroup) *ServerController {
	a := &ServerController{}
	a.initRouter(g)
	return a
}

func (a *ServerController) initRouter(g *gin.RouterGroup) {
	g = g.Group("/server")

	g.POST("/restartXrayService", a.restartXrayService)
	g.POST("/getNewX25519Cert", a.getNewX25519Cert)
}

func (a *ServerController) restartXrayService(c *gin.Context) {
	err := a.serverService.RestartXrayService()
	if err != nil {
		jsonMsg(c, "", err)
		return
	}
	jsonMsg(c, "Xray restarted", err)
}

func (a *ServerController) getNewX25519Cert(c *gin.Context) {
	cert, err := a.serverService.GetNewX25519Cert()
	if err != nil {
		jsonMsg(c, "get x25519 certificate", err)
		return
	}
	jsonObj(c, cert, nil)
}
