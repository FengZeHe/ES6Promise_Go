package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// 定义一个简单的数据结构（模拟用户数据）
type User struct {
	ID   string `json:"id" form:"id"` // 支持JSON和表单参数
	Name string `json:"name" form:"name"`
	Age  int    `json:"age" form:"age"`
}

func main() {
	r := gin.Default()
	r.Use(corsMiddleware())
	r.GET("/hi", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello",
		})
	})

	if err := r.Run(":8080"); err != nil {
		panic(err)
	}

}

func corsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		c.Writer.Header().Set("Access-Control-Expose-Headers", "Content-Length")
		c.Writer.Header().Set("Access-Control-Max-Age", "86400")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}

		c.Next()
	}
}
