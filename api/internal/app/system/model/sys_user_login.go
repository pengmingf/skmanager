/*
* @desc:登录日志
* @company:舒可到家健康管理有限公司
* @Author: pmf
* @Date:   2022/3/8 11:43
 */

package model

// LoginLogParams 登录日志写入参数
type LoginLogParams struct {
	Status    int
	Username  string
	Ip        string
	UserAgent string
	Msg       string
	Module    string
}
