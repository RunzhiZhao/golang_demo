package router

import (
	"github.com/gin-gonic/gin"
	"golang_demo/handler"
	"net/http"
)

func StartRouter() {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "It Works")
	})

	// 1. 创建插入demo_order
	router.POST("/createOrder", handler.CreateOrder)

	// 2. 更新 demo_order （amount、status、file_url）
	router.POST("/updateOrder", handler.UpdateOrder)

	// 3. 获取 demo_order 详情
	router.GET("/getOrderInfo", handler.GetOrderInfo)

	// 4. 获取 demo_order 列表 （需要包含： 模糊查找、根据创建时间，金额排序）
	router.POST("/getOrders", handler.GetOrders)

	// 5. gin上传文件， 更新fileUrl
	// Set a lower memory limit for multipart forms (default is 32 MiB)
	router.MaxMultipartMemory = 8 << 20 // 8 MiB
	//router.Static("/", "./public")
	router.POST("/uploadFile", handler.UploadFile)

	// 6. 下载文件
	router.GET("downloadFile", handler.DownloadFile)

	// 7. order表转excel并下载
	router.GET("/getOrderExcel", handler.ExportOrderExcel)

	router.Run(":8080")
}
