package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/piggona/fundingsView/api/dbops"
)

func GetDetail(c *gin.Context) {
	uuid := c.Param("uuid")
	resp, err := dbops.GetAwardIDDetail(uuid, "nsf_test")
	if err != nil {
		fmt.Errorf("Error get detail: %s", err)
		return
	}
	c.JSON(http.StatusOK, resp)
	return
}

func GetCopTree(c *gin.Context) {
	return
}

func GetWordTree(c *gin.Context) {
	return
}

func GetSimilar(c *gin.Context) {
	return
}
