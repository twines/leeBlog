package admins

import (
	"github.com/gin-gonic/gin"
	"leeblog.com/app/http/response"
	"net/http"
	"time"
)

// @Tags admin
// @Summary 订单列表
// @Description # 请求参数
// @Description | 参数 | 说明 | 备注 |
// @Description | :-----| :----- | :----- |
// @Description |page|页码||
// @Description # 返回参数参数
// @Description | 参数 | 说明 | 备注 |
// @Description | :-----| :----- | :----- |
// @Description |code|状态码||
// @Description |message|返回的信息||
// @Description |data|返回的具体数据||
// @Produce  json
// @Security bearer
// @Param page query string false "页码"
// @Success 200 {object} response.Response "{"code":200,"data":{},"msg":"ok"}"
// @Router /admin/v1/order/list [get]
func OrderIndex(c *gin.Context) {
	m := map[string]interface{}{}
	m["total"] = 20
	var orderSlice []interface{}
	for i := 0; i < 20; i++ {
		order := map[string]interface{}{}
		order["order_no"] = "111"
		order["timestamp"] = time.Now()
		order["username"] = "111"
		order["price"] = "111"
		order["price"] = "111"
		if i%2 == 0 {
			order["status"] = "pending"
		} else {
			order["status"] = "success"
		}
		orderSlice = append(orderSlice, order)
	}
	m["items"] = orderSlice

	r := response.Response{Data: m, Code: 20000}
	c.JSON(http.StatusOK, r)
}
