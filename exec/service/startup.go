package service

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/yuansudong/http_dns/flags"
)

// Startup 用于启动一个服务
func Startup(iGlobalFlag *flags.GlobalFlag,
	iLocalFlag *flags.LocalServiceFlag,
) {

	router := gin.Default()

	router.GET("/query", func(c *gin.Context) {
		// Request 请求
		type Request struct {
			Domain string `json:"domain"`
		}
		// Response 响应
		type Response struct {
			IP string `json:"ip"`
		}
		rsp := new(Response)
		rsp.IP = "14.215.177.39"
		log.Println("host:", c.GetHeader("host"), c.GetHeader("Host"))
		log.Printf("%+v\n", c.Request.Host)
		c.JSON(200, rsp)
	})
	router.Run(":5001")
}
