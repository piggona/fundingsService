package main

import (
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/piggona/fundingsView/api/defs"
)

func sendErrorResponse(c *gin.Context, errResp defs.ErrorResponse) {
	c.JSON(errResp.HTTPSc, gin.H{"data": &errResp.Error})
}

func sendNormalResponse(w http.ResponseWriter, resp string, sc int) {
	w.WriteHeader(sc)
	io.WriteString(w, resp)
}
