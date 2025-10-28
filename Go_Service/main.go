package main

import (
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},                                                 // 允许所有来源（生产环境需指定具体域名）
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},  // 允许的请求方法
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"}, // 允许的请求头
		ExposeHeaders:    []string{"Content-Length"},                                    // 允许前端读取的响应头
		AllowCredentials: true,                                                          // 允许携带Cookie（跨域请求时）
		MaxAge:           12 * time.Hour,                                                // 预检请求的缓存时间（12小时）
	}))

	r.GET("/hi", func(c *gin.Context) {
		random := rand.Intn(3)
		if random == 0 {
			log.Println("返回失败")
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "ERROR",
			})
		} else {
			log.Println("返回成功")
			c.JSON(http.StatusOK, gin.H{
				"message": "Hello",
			})
		}
	})

	if err := r.Run(":8080"); err != nil {
		panic(err)
	}

}
