/*
* @desc:验证码参数
* @company:舒可到家健康管理有限公司
* @Author: pmf
* @Date:   2022/3/2 17:47
 */

package common

import "github.com/gogf/gf/v2/frame/g"

type CaptchaReq struct {
	g.Meta `path:"/get" tags:"验证码" method:"get" summary:"获取验证码"`
}
type CaptchaRes struct {
	g.Meta `mime:"application/json"`
	Key    string `json:"key"`
	Img    string `json:"img"`
}
