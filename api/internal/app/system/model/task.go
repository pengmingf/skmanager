/*
* @desc:定时任务
* @company:舒可到家健康管理有限公司
* @Author: pmf<291641454@qq.com>
* @Date:   2023/1/13 17:47
 */

package model

import "context"

type TimeTask struct {
	FuncName string
	Param    []string
	Run      func(ctx context.Context)
}
