/*
* @desc:错误处理
* @company:舒可到家健康管理有限公司
* @Author: pmf
* @Date:   2022/3/2 14:53
 */

package liberr

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"
)

func ErrIsNil(ctx context.Context, err error, msg ...string) {
	if !g.IsNil(err) {
		if len(msg) > 0 {
			g.Log().Error(ctx, err.Error())
			panic(msg[0])
		} else {
			panic(err.Error())
		}
	}
}

func ValueIsNil(value interface{}, msg string) {
	if g.IsNil(value) {
		panic(msg)
	}
}
