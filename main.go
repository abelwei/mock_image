package main

import (
	"github.com/gin-gonic/gin"
	"mock_image/draw_image"
	"net/http"
)

func main() {
	r := gin.Default()
	r.GET("/draw", func(c *gin.Context) {
		//http://localhost:8080/draw?dsl=rect,w=500,h=200,color=ff0000
		drawForm := c.Query("dsl")
		err, bt := draw_image.NewDrawPattern().Parse(drawForm).Response()
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.Writer.Write(bt)
	})
	r.Run()
}
