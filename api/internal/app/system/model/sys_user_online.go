/*
* @desc:用户在线状态
* @company:舒可到家健康管理有限公司
* @Author: pmf<291641454@qq.com>
* @Date:   2023/1/10 15:08
 */

package model

// SysUserOnlineParams 用户在线状态写入参数
type SysUserOnlineParams struct {
	UserAgent string
	Uuid      string
	Token     string
	Username  string
	Ip        string
}
