package v1

import (
	"fmt"
	"gin-vue-admin/global/response"
	"gin-vue-admin/model/request"
	resp "gin-vue-admin/model/response"
	"gin-vue-admin/service"
	"github.com/gin-gonic/gin"
)

type stockTop int

// @Tags Stock
// @Summary 获取Top列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.StockTopListReq true "获取Top列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /stock/top/list [post]
func (*stockTop) StockTopList(c *gin.Context) {
	var req request.StockTopListReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.FailWithMessage(fmt.Sprintf("参数错误，%v", err), c)
		return
	}

	stockList, err := service.StockTopService.TopList()
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("获取Top失败，%v", err), c)
	} else {
		response.OkWithData(resp.PageResult{
			List:     stockList,
			Total:    len(stockList),
			Page:     1,
			PageSize: len(stockList),
		}, c)
	}
}
