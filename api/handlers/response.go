package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/piggona/fundingsView/api/defs"
)

func sendErrorResponse(c *gin.Context, errResp defs.ErrorResponse) {
	c.JSON(errResp.HTTPSc, gin.H{"data": &errResp.Error})
}

func sendNormalResponse(c *gin.Context, resp defs.NormalResp) {
	c.JSON(resp.HttpSc, gin.H{"data": resp.Resp})
}
