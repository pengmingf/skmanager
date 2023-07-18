/*
* @desc:system base controller
* @company:舒可到家健康管理有限公司
* @Author: pmf
* @Date:   2022/3/4 18:12
 */

package controller

import (
	"github.com/gogf/gf/v2/net/ghttp"
	commonController "github.com/tiger1103/gfast/v3/internal/app/common/controller"
)

type BaseController struct {
	commonController.BaseController
}

// Init 自动执行的初始化方法
func (c *BaseController) Init(r *ghttp.Request) {
	c.BaseController.Init(r)
}
