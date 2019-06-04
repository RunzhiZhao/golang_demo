package handler

import (
	"github.com/gin-gonic/gin"
	"golang_demo/model"
	"golang_demo/service"
	"net/http"
)

// 1. 创建插入demo_order
func CreateOrder(c *gin.Context) {

	var req model.CreateOrderReq
	// 参数格式检测
	if err := c.BindJSON(&req); err != nil {
		response := model.ResponseData{Code: 0, Status: "error", Message:"参数错误", Data: ""}
		c.JSON(http.StatusOK, response)
		return
	}

	if err := service.CreateOrderService(&req); err != nil {
		response := model.ResponseData{Code: 0, Status: "error", Message:err.Error(), Data: ""}
		c.JSON(http.StatusOK, response)
		return
	} else  {
		response := model.ResponseData{Code: 1, Status: "success", Message:"操作成功", Data: ""}
		c.JSON(http.StatusOK, response)
	}
}


// 2. 更新 demo_order （amount、status、file_url）
func UpdateOrder(c *gin.Context)  {
	var req model.UpdateOrderReq
	if err := c.BindJSON(&req); err != nil {
		response := model.ResponseData{Code: 0, Status: "error", Message:"参数错误", Data: ""}
		c.JSON(http.StatusBadRequest, response)
		return
	}

	if err := service.UpdateOrderService(&req); err != nil {
		response := model.ResponseData{Code: 0, Status: "error", Message:err.Error(), Data: ""}
		c.JSON(http.StatusOK, response)
		return
	} else  {
		response := model.ResponseData{Code: 1, Status: "success", Message:"操作成功", Data: ""}
		c.JSON(http.StatusOK, response)
	}
}


// 3. 获取 demo_order 详情
func GetOrderInfo(c *gin.Context)  {
	var req model.GetOrderInfoReq
	if err := c.BindJSON(&req); err != nil {
		response := model.ResponseData{Code: 0, Status: "error", Message:"参数错误", Data: ""}
		c.JSON(http.StatusOK, response)
		return
	}

	if order := service.GetOrderInfoService(&req); order != nil {
		response := model.ResponseData{Code: 1, Status: "success", Message:"操作成功", Data: order}
		c.JSON(http.StatusOK, response)
		return
	} else  {
		response := model.ResponseData{Code: 0, Status: "error", Message:"订单不存在", Data: ""}
		c.JSON(http.StatusOK, response)
	}
}

// 4. 获取 demo_order 列表 （需要包含： 模糊查找、根据创建时间，金额排序）
func GetOrders(c *gin.Context)  {
	var req model.GetOrderInfoReq

	if err := c.BindJSON(&req); err != nil {
		response := model.ResponseData{Code: 0, Status: "error", Message:"参数错误", Data: ""}
		c.JSON(http.StatusOK, response)
		return
	}


}