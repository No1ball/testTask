package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h *Handler) SendEmail(c *gin.Context) {
	arr, _ := c.GetQueryArray("arr[]")
	htmlTemplates := c.Query("form")
	id1, _ := strconv.Atoi(htmlTemplates)
	c.JSON(http.StatusOK, gin.H{
		"id1":  arr[0],
		"form": id1,
	})
}
