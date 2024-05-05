package main

import (
	"github.com/gin-gonic/gin"
	"github.com/tonymelek/go_first/initialisers"
)

func init() {
	initialisers.LoadEnv();
}
func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.String(200, "Tony is testing GOLANG")
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for default)
}
